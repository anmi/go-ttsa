package helpers

import (
	"context"
	"database/sql"
	"time"

	ttssqlc "github.com/anmi/go-ttsa/db"
)

func CreateRootTask(
	ctx context.Context,
	q *ttssqlc.Queries,
	user_id int64,
) (int64, error) {

	root, root_err := q.CreateTask(ctx, ttssqlc.CreateTaskParams{
		Title:       "",
		Description: "",
		Result:      "",
		CreatedAt:   time.Now(),
		CreatedBy:   user_id,
	})

	if root_err != nil {
		return 0, root_err
	}

	return root.ID, q.UpdateRootTaskId(ctx, ttssqlc.UpdateRootTaskIdParams{
		ID:         user_id,
		RootTaskID: sql.NullInt64{Valid: true, Int64: root.ID},
	})
}
