package helpers

import (
	"context"

	queries "github.com/anmi/go-ttsa/db"
	"github.com/anmi/go-ttsa/utils"
)

func GetSession(ctx context.Context, q queries.Queries) (queries.GetSessionRow, error) {
	req := utils.RequestFromContext(ctx)
	token, error := req.Cookie("GSESSIONID")
	if error != nil {
		return queries.GetSessionRow{}, error
	}

	session, err := q.GetSession(ctx, token.Value)
	if error != nil {
		return queries.GetSessionRow{}, err
	}

	return session, nil
}
