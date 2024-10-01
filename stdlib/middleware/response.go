package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	apperr "github.com/linggaaskaedo/go-rocks/stdlib/errors"
	"github.com/linggaaskaedo/go-rocks/stdlib/preference"
)

type HTTPErrResp struct {
	Meta Meta `json:"metadata"`
}

type Meta struct {
	Path       string           `json:"path"`
	StatusCode int              `json:"status_code"`
	Status     string           `json:"status"`
	Message    string           `json:"message"`
	Error      *apperr.AppError `json:"error,omitempty" swaggertype:"primitive,object"`
	Timestamp  string           `json:"timestamp"`
}

func (m *middleware) httpRespError(c *gin.Context, err error) *HTTPErrResp {
	ctx := c.Request.Context()

	zerolog.Ctx(ctx).Err(err).Send()

	lang := preference.LangID

	if c.Request.Header[preference.AppLang] != nil && c.Request.Header[preference.AppLang][0] == preference.LangEN {
		lang = preference.LangEN
	}

	statusCode, displayError := apperr.Compile(apperr.COMMON, err, lang, true)
	statusStr := http.StatusText(statusCode)

	jsonErrResp := &HTTPErrResp{
		Meta: Meta{
			Path:       c.Request.URL.Path,
			StatusCode: statusCode,
			Status:     statusStr,
			Message:    fmt.Sprintf("%s %s [%d] %s", c.Request.Method, c.Request.RequestURI, statusCode, http.StatusText(statusCode)),
			Error:      &displayError,
			Timestamp:  time.Now().Format(time.RFC3339),
		},
	}

	return jsonErrResp
}
