package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/linggaaskaedo/go-rocks/src/business/entity"
	"github.com/linggaaskaedo/go-rocks/src/common"
	commonerr "github.com/linggaaskaedo/go-rocks/stdlib/errors/common"
	x "github.com/linggaaskaedo/go-rocks/stdlib/errors/entity"
	"github.com/linggaaskaedo/go-rocks/stdlib/preference"
	"github.com/linggaaskaedo/go-rocks/stdlib/preference/sqlx"
)

// CreateDivision		godoc
//
//	@Summary		Create division
//	@Description	Endpoint for creating a new division
//	@Tags			Division
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			Authorization	header		string					true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Param			data			body		DivisionCreateRequest	true	"Payload create division"
//	@Success		201				{object}	HTTPDivisionResp
//	@Failure		400				{object}	HTTPErrResp
//	@Failure		401				{object}	HTTPErrResp
//	@Failure		429				{object}	HTTPErrResp
//	@Failure		500				{object}	HTTPErrResp
//	@Router			/division [post]
func (e *rest) CreateDivision(c *gin.Context) {
	ctx := c.Request.Context()

	var requestBody DivisionCreateRequest

	if err := c.ShouldBind(&requestBody); err != nil {
		e.httpRespError(c, x.WrapWithCode(err, commonerr.CodeHTTPUnmarshal, "unmarshall_err"))
		return
	}

	if requestBody.Data.Division == nil {
		e.httpRespError(c, x.NewWithCode(commonerr.CodeHTTPBadRequest, "invalid_payload"))
		return
	}

	data := requestBody.Data.Division

	validationErr := data.Validate()
	if validationErr != nil {
		e.httpRespError(c, x.WrapWithCode(validationErr, commonerr.CodeHTTPBadRequest, "validation_error"))
		return
	}

	divisionEntity := entity.Division{
		Name:      data.Name,
		CreatedAt: sqlx.NullTime{Valid: true, Time: time.Now()},
	}

	result, err := e.uc.Division.CreateDivision(ctx, divisionEntity)
	if err != nil {
		e.httpRespError(c, err)
		return
	}

	e.httpRespSuccess(c, http.StatusCreated, result, nil)
}

// GetDivisionByID		godoc
//
//	@Summary		Get division by ID
//	@Description	Endpoint for get a division with ID
//	@Tags			Division
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			Authorization	header		string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Param			divisionID		path		string	true	"Division ID"
//	@Success		200				{object}	HTTPDivisionResp
//	@Failure		400				{object}	HTTPErrResp
//	@Failure		401				{object}	HTTPErrResp
//	@Failure		500				{object}	HTTPErrResp
//	@Router			/division/{divisionID} [get]
func (e *rest) GetDivisionByID(c *gin.Context) {
	ctx := c.Request.Context()

	varID := c.Param("divisionID")

	divisionID, err := common.MixerDecode(varID)
	if err != nil {
		e.httpRespError(c, x.Wrap(err, "decode_division_id"))
		return
	}

	result, err := e.uc.Division.GetDivisioByID(ctx, divisionID)
	if err != nil {
		e.httpRespError(c, err)
		return
	}

	e.httpRespSuccess(c, http.StatusOK, result, nil)
}

// GetDivision		godoc
//
//	@Summary		Get list of division based on query params
//	@Description	Endpoint for get division with param
//	@Tags			Division
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			Authorization	header		string		true	"Insert your access token"		default(Bearer <Add access token here>)
//	@Param			Cache-Control	header		string		false	"Request cache control"			Enums(must-revalidate, must-db-revalidate)
//	@Param			public_id		query		[]string	false	"Search by Division Public ID"	collectionFormat(multi)
//	@Param			name			query		string		false	"Search by Division name. use % for wildcard search."
//	@Param			is_deleted		query		string		false	"Search by deleted status"			Enums(false,true)
//	@Param			sort_by			query		string		false	"Sort result by these attributes"	Enums(id, -id, name, -name, created_at, -created_at)	default(-created_at)
//	@Param			page			query		string		false	" "
//	@Param			limit			query		string		false	" "
//	@Success		200				{object}	HTTPDivisionsResp
//	@Failure		400				{object}	HTTPErrResp
//	@Failure		401				{object}	HTTPErrResp
//	@Failure		500				{object}	HTTPErrResp
//	@Router			/division [get]
func (e *rest) GetDivision(c *gin.Context) {
	ctx := c.Request.Context()

	var param entity.DivisionParam

	if err := e.param.Decode(&param, c.Request.URL.Query()); err != nil {
		e.httpRespError(c, x.WrapWithCode(err, commonerr.CodeHTTPBadRequest, "decode_query_param"))
		return
	}

	if len(param.PublicID) != 0 {
		var ids []int64

		for _, s := range param.PublicID {
			id, err := common.MixerDecode(s)
			if err != nil {
				e.httpRespError(c, x.WrapWithCode(err, commonerr.CodeHTTPBadRequest, "decode_division_id"))
				return
			}

			ids = append(ids, id)
		}

		param.ID = ids
	}

	if len(param.SortBy) == 0 {
		param.SortBy = append(param.SortBy, "-created_at")
	}

	zerolog.Ctx(ctx).Debug().Any("PARAM", param).Send()

	var cacheControl entity.CacheControl

	if c.Request.Header[preference.CacheControl] != nil && c.Request.Header[preference.CacheControl][0] == preference.CacheMustRevalidate {
		cacheControl.MustRevalidate = true
	}

	results, pagination, err := e.uc.Division.GetDivision(ctx, cacheControl, param)
	if err != nil {
		e.httpRespError(c, err)
		return
	}

	e.httpRespSuccess(c, http.StatusOK, results, &pagination)
}
