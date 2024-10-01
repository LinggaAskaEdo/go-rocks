package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/linggaaskaedo/go-rocks/src/business/dto"
	"github.com/linggaaskaedo/go-rocks/src/business/entity"
	apperr "github.com/linggaaskaedo/go-rocks/stdlib/errors"
	x "github.com/linggaaskaedo/go-rocks/stdlib/errors/entity"
	"github.com/linggaaskaedo/go-rocks/stdlib/preference"
)

func (e *rest) httpRespSuccess(c *gin.Context, statusCode int, resp interface{}, p *entity.Pagination) {
	var raw interface{}

	meta := dto.Meta{
		Path:       c.Request.URL.Path,
		StatusCode: statusCode,
		Status:     http.StatusText(statusCode),
		Message:    fmt.Sprintf("%s %s [%d] %s", c.Request.Method, c.Request.RequestURI, statusCode, http.StatusText(statusCode)),
		Error:      nil,
		Timestamp:  time.Now().Format(time.RFC3339),
	}

	switch data := resp.(type) {
	case nil:
		httpResp := &HTTPEmptyResp{
			Meta: meta,
		}
		raw = httpResp

	case dto.UserDTO:
		httpResp := &HTTPUserResp{
			Meta: meta,
			Data: UserData{
				User: &data,
			},
		}
		raw = httpResp

	case dto.UserLoginDTO:
		httpResp := &HTTPUserLoginResp{
			Meta: meta,
			Data: UserLoginData{
				Token: &data,
			},
		}
		raw = httpResp

	case dto.UserLogoutDTO:
		httpResp := &HTTPUserLogoutResp{
			Meta: meta,
			Data: UserLogoutData{
				Token: &data,
			},
		}
		raw = httpResp

	case dto.DivisionDTO:
		httpResp := &HTTPDivisionResp{
			Meta: meta,
			Data: DivisionData{
				Division: &data,
			},
		}
		raw = httpResp

	case []dto.DivisionDTO:
		httpResp := &HTTPDivisionsResp{
			Meta: meta,
			Data: DivisionsData{
				Divisions: data,
			},
			Pagination: p,
		}
		raw = httpResp

	case dto.EmployeeDTO:
		httpResp := &HTTPEmployeeResp{
			Meta: meta,
			Data: EmployeeData{
				Employee: &data,
			},
		}
		raw = httpResp

	case []dto.EmployeeDTO:
		httpResp := &HTTPEmployeesResp{
			Meta: meta,
			Data: EmployeesData{
				Employees: data,
			},
			Pagination: p,
		}
		raw = httpResp

	case dto.KCEmployeeDTO:
		httpResp := &HTTPKCEmployeeResp{
			Meta: meta,
			Data: KCEmployeeData{
				Employee: &data,
			},
		}
		raw = httpResp

	default:
		e.httpRespError(c, x.New("Invalid response type"))
		return
	}

	c.JSON(statusCode, raw)
}

func (e *rest) httpRespError(c *gin.Context, err error) {
	lang := preference.LangID

	if c.Request.Header[preference.AppLang] != nil && c.Request.Header[preference.AppLang][0] == preference.LangEN {
		lang = preference.LangEN
	}

	statusCode, displayError := apperr.Compile(apperr.COMMON, err, lang, true)
	statusStr := http.StatusText(statusCode)

	jsonErrResp := &HTTPErrResp{
		Meta: dto.Meta{
			Path:       c.Request.URL.Path,
			StatusCode: statusCode,
			Status:     statusStr,
			Message:    fmt.Sprintf("%s %s [%d] %s", c.Request.Method, c.Request.RequestURI, statusCode, http.StatusText(statusCode)),
			Error:      &displayError,
			Timestamp:  time.Now().Format(time.RFC3339),
		},
	}

	c.JSON(statusCode, jsonErrResp)
}
