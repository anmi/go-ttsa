package mapping

import (
	"github.com/anmi/go-ttsa/api"
	ttssqlc "github.com/anmi/go-ttsa/db"
	helpers "github.com/anmi/go-ttsa/service/utils"
)

func SqlTaskToApi(task ttssqlc.Task) api.Task {
	return api.Task{
		ID:          int(task.ID),
		Title:       task.Title,
		Description: task.Description,
		Result:      task.Result,
		CreatedAt:   task.CreatedAt.String(),
		DoneAt:      helpers.NilableTimeToApi(task.DoneAt),
	}
}
