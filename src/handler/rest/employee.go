package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/linggaaskaedo/go-rocks/src/business/entity"
	"github.com/linggaaskaedo/go-rocks/src/common"
	commonerr "github.com/linggaaskaedo/go-rocks/stdlib/errors/common"
	x "github.com/linggaaskaedo/go-rocks/stdlib/errors/entity"
	"github.com/linggaaskaedo/go-rocks/stdlib/preference"
)

// GetEmployeeByID		godoc
//
//	@Summary		Get employee by ID
//	@Description	Endpoint for get a employee with ID
//	@Tags			Employee
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			Authorization	header		string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Param			employeeID		path		string	true	"Employee ID"
//	@Success		200				{object}	HTTPEmployeeResp
//	@Failure		400				{object}	HTTPErrResp
//	@Failure		401				{object}	HTTPErrResp
//	@Failure		500				{object}	HTTPErrResp
//	@Router			/employee/{employeeID} [get]
func (e *rest) GetEmployeeByID(c *gin.Context) {
	ctx := c.Request.Context()

	varID := c.Param("employeeID")

	employeeID, err := common.MixerDecode(varID)
	if err != nil {
		e.httpRespError(c, x.Wrap(err, "decode_employee_id"))
		return
	}

	result, err := e.uc.Employee.GetEmployeeByID(ctx, employeeID)
	if err != nil {
		e.httpRespError(c, err)
		return
	}

	e.httpRespSuccess(c, http.StatusOK, result, nil)
}

// GetEmployee		godoc
//
//	@Summary		Get list of employee based on query params
//	@Description	Endpoint for get a employee with param
//	@Tags			Employee
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			Authorization	header		string		true	"Insert your access token"		default(Bearer <Add access token here>)
//	@Param			Cache-Control	header		string		false	"Request cache control"			Enums(must-revalidate, must-db-revalidate)
//	@Param			public_id		query		[]string	false	"Search by Employee Public ID"	collectionFormat(multi)
//	@Param			first_name		query		string		false	"Search by Employee first name"
//	@Param			last_name		query		string		false	"Search by Employee last name"
//	@Param			gender			query		string		false	"Search by Employee gender"			Enums(M, F)
//	@Param			sort_by			query		string		false	"Sort result by these attributes"	Enums(id, -id, firstname, -firstname, lastname, -lastname, gender, -gender)	default(id)
//	@Param			page			query		string		false	" "
//	@Param			limit			query		string		false	" "
//	@Success		200				{object}	HTTPEmployeesResp
//	@Failure		400				{object}	HTTPErrResp
//	@Failure		401				{object}	HTTPErrResp
//	@Failure		500				{object}	HTTPErrResp
//	@Router			/employee [get]
func (e *rest) GetEmployee(c *gin.Context) {
	ctx := c.Request.Context()

	var param entity.EmployeeParam

	if err := e.param.Decode(&param, c.Request.URL.Query()); err != nil {
		e.httpRespError(c, x.WrapWithCode(err, commonerr.CodeHTTPBadRequest, "decode_query_param"))
		return
	}

	if len(param.PublicID) != 0 {
		var ids []int64

		for _, s := range param.PublicID {
			id, err := common.MixerDecode(s)
			if err != nil {
				e.httpRespError(c, x.WrapWithCode(err, commonerr.CodeHTTPBadRequest, "decode_employee_id"))
				return
			}

			ids = append(ids, id)
		}

		param.ID = ids
	}

	if len(param.SortBy) == 0 {
		param.SortBy = append(param.SortBy, "id")
	}

	zerolog.Ctx(ctx).Debug().Any("PARAM", param).Send()

	var cacheControl entity.CacheControl

	if c.Request.Header[preference.CacheControl] != nil && c.Request.Header[preference.CacheControl][0] == preference.CacheMustRevalidate {
		cacheControl.MustRevalidate = true
	}

	results, pagination, err := e.uc.Employee.GetEmployee(ctx, cacheControl, param)
	if err != nil {
		e.httpRespError(c, err)
		return
	}

	e.httpRespSuccess(c, http.StatusOK, results, &pagination)
}
