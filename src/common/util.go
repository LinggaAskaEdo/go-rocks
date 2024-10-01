package common

import (
	"github.com/linggaaskaedo/go-rocks/stdlib/preference/sqlx"
)

func StringTime(data sqlx.NullTime) string {
	var result string

	if data.Valid {
		result = data.Time.Format("02-01-2006")
	} else {
		result = "-"
	}

	return result
}
