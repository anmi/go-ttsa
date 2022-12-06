package helpers

import (
	"database/sql"

	api "github.com/anmi/go-ttsa/api"
)

func NilableTimeToApi(time sql.NullTime) api.OptString {
	if time.Valid {
		return api.OptString{Set: true, Value: time.Time.String()}
	}

	return api.OptString{Set: false}
}
