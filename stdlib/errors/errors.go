package errors

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/linggaaskaedo/go-rocks/stdlib/errors/common"
	"github.com/linggaaskaedo/go-rocks/stdlib/errors/entity"
	"github.com/linggaaskaedo/go-rocks/stdlib/preference"
)

const (
	COMMON ServiceType = 1
)

var svcError map[ServiceType]entity.ErrorMessage

type ServiceType int

type AppError struct {
	Code       entity.Code `json:"code"`
	Message    string      `json:"message"`
	DebugError *string     `json:"debug,omitempty"`
	sys        error
}

func init() {
	svcError = map[ServiceType]entity.ErrorMessage{
		COMMON: common.ErrorMessages,
	}
}

func Compile(service ServiceType, err error, lang string, debugMode bool) (int, AppError) {
	var debugErr *string

	if debugMode {
		errStr := err.Error()
		if len(errStr) > 0 {
			debugErr = &errStr
		}
	}

	code := entity.ErrCode(err)

	if errMessage, ok := svcError[COMMON][code]; ok {
		msg := errMessage.ID
		if lang == preference.LangEN {
			msg = errMessage.EN
		}

		return errMessage.StatusCode, AppError{
			Code:       code,
			Message:    msg,
			sys:        err,
			DebugError: debugErr,
		}
	}

	if errMessages, ok := svcError[service]; ok {
		if errMessage, ok := errMessages[code]; ok {
			msg := errMessage.ID
			if lang == preference.LangEN {
				msg = errMessage.EN
			}

			if errMessage.HasAnnotation {
				args := fmt.Sprintf("%q", err.Error())
				if start, end := strings.LastIndex(args, `{{`), strings.LastIndex(args, `}}`); start > -1 && end > -1 {
					args = strings.TrimSpace(args[start+2 : end])
					msg = fmt.Sprintf(msg, args)
				} else {
					index := strings.Index(args, `\n`)
					if index > 0 {
						args = strings.TrimSpace(args[1:index])
					}

					msg = fmt.Sprintf(msg, args)
				}
			}

			if code == common.CodeHTTPValidatorError {
				if err.Error() != "" {
					msg = strings.Split(err.Error(), "\n ---")[0]
				}
			}

			return errMessage.StatusCode, AppError{
				Code:       code,
				Message:    msg,
				sys:        err,
				DebugError: debugErr,
			}
		}

		return http.StatusInternalServerError, AppError{
			Code:       code,
			Message:    "error message not defined!",
			sys:        err,
			DebugError: debugErr,
		}
	}

	return http.StatusInternalServerError, AppError{
		Code:       code,
		Message:    "service error not defined!",
		sys:        err,
		DebugError: debugErr,
	}
}
