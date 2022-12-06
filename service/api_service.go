package service

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	api "github.com/anmi/go-ttsa/api"
	q "github.com/anmi/go-ttsa/db"
	helpers "github.com/anmi/go-ttsa/service/utils"
	utils "github.com/anmi/go-ttsa/utils"
	"golang.org/x/crypto/bcrypt"
)

type TTApiService struct {
	Queries *q.Queries
}

func (s *TTApiService) GetCurrentUser(ctx context.Context) (api.GetCurrentUserRes, error) {
	session, err := helpers.GetSession(ctx, *s.Queries)

	if err != nil {
		return ServiceErrors.Unauthorized(), nil
	}

	RootID := api.OptInt{Set: false}
	if session.RootTaskID.Valid {
		RootID = api.NewOptInt(int(session.RootTaskID.Int64))
	}

	return &api.CurrentUser{
		Username: session.Username,
		Email:    session.Email.String,
		RootID:   RootID,
	}, nil
}

func (s *TTApiService) LogOut(ctx context.Context) (api.LogOutOK, error) {
	req := utils.RequestFromContext(ctx)
	token, error := req.Cookie("GSESSIONID")
	if error != nil {
		return api.LogOutOK{}, error
	}
	err := s.Queries.DeleteSession(ctx, token.Value)
	if err != nil {
		return api.LogOutOK{}, err
	}
	return api.LogOutOK{}, nil
}

func (s *TTApiService) SignIn(ctx context.Context, req api.OptSignInForm) (api.SignInRes, error) {
	user, err := s.Queries.GetUserByName(ctx, req.Value.Username)

	// Couldn't get User from DB
	if err != nil {
		return &api.SignInWrongUsernameResponse{
			Errorrr: api.OptSignInWrongUsernameResponseErrorrr{Value: "WrongUsername"},
		}, nil
	}
	// No password has been set for user
	if !user.Password.Valid || !CheckPasswordHash(req.Value.Password, user.Password.String) {
		return &api.SignInWrongUsernameResponse{
			Errorrr: api.OptSignInWrongUsernameResponseErrorrr{Value: "WrongUsername"},
		}, nil
	}

	token := helpers.GenerateToken()

	fmt.Println("Token %V", token)

	// Create session
	session, err := s.Queries.CreateSession(ctx, q.CreateSessionParams{
		UserID:    user.ID,
		Token:     token,
		CreatedAt: time.Now(),
	})

	if err != nil {
		return &api.SignInResponse{}, err
	}

	cookie := http.Cookie{
		Name:  "GSESSIONID",
		Value: session.Token,
		Path:  "/",
	}
	cookie.Expires = time.Now().Add(time.Hour)
	res_wr := utils.ResponseWriterFromContext(ctx)
	http.SetCookie(*res_wr, &cookie)

	return &api.SignInResponse{Ok: "success"}, nil
}

type errorUserExist struct{}

func (s errorUserExist) Error() string {
	return "User already exist"
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *TTApiService) SignUp(ctx context.Context, req api.OptSignUpForm) (api.SignUpOK, error) {
	_, err := s.Queries.GetUserByName(ctx, req.Value.Username)

	if err == nil {
		return api.SignUpOK{}, errorUserExist{}
	}

	if err != sql.ErrNoRows {
		return api.SignUpOK{}, err
	}

	hash, _ := HashPassword(req.Value.Password)

	created, err := s.Queries.CreateUser(ctx, q.CreateUserParams{
		Username:  req.Value.Username,
		Password:  sql.NullString{String: hash, Valid: true},
		Email:     sql.NullString{String: req.Value.Email, Valid: true},
		CreatedAt: time.Now(),
	})

	if err != nil {
		return api.SignUpOK{}, err
	}

	fmt.Printf("Created user %v\n", created.Username)

	return api.SignUpOK{}, nil
}
