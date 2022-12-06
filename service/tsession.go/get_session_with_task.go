package tsession

import (
	"context"

	queries "github.com/anmi/go-ttsa/db"
	"github.com/anmi/go-ttsa/utils"
)

type SessionError struct {
	message string
}

func (se SessionError) Error() string {
	return se.message
}

type TSession struct {
	UserID     int64
	Authorized bool
	queries    *queries.Queries
	ctx        context.Context
}

var SessionUnauthorizedError = SessionError{
	message: "SessionUnauthorizedError",
}

var SessionForbiddenError = SessionError{
	message: "SessionForbiddenError",
}

func GetTSession(ctx context.Context, q *queries.Queries) (*TSession, error) {
	req := utils.RequestFromContext(ctx)

	token, token_err := req.Cookie("GSESSIONID")
	if token_err != nil {
		return &TSession{}, SessionUnauthorizedError
	}

	session_response, session_err := q.GetSession(ctx, token.Value)
	if session_err != nil {
		return &TSession{}, SessionUnauthorizedError
	}

	session := TSession{
		UserID:     session_response.UserID,
		Authorized: token_err == nil && session_err == nil,
		queries:    q,
		ctx:        ctx,
	}

	return &session, nil
}

func (s TSession) GetTask(id int64) (queries.Task, error) {
	if !s.Authorized {
		return queries.Task{}, SessionUnauthorizedError
	}
	task, err := s.queries.GetTask(s.ctx, id)

	if err != nil {
		return queries.Task{}, err
	}

	if task.CreatedBy != s.UserID {
		return queries.Task{}, SessionForbiddenError
	}

	return task, nil
}
