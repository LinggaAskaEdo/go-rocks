package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/linggaaskaedo/go-rocks/src/business/dto"
	commonerr "github.com/linggaaskaedo/go-rocks/stdlib/errors/common"
	x "github.com/linggaaskaedo/go-rocks/stdlib/errors/entity"
)

// KC - Login 	godoc
//
//	@Summary		Keycloak Login
//	@Description	Endpoint for generating access token for user via keycloak
//	@Tags			KC
//	@Accept			json
//	@Produce		json
//	@Param			data	body		UserLoginRequest	true	"Payload user login"
//	@Success		200		{object}	HTTPUserLoginResp
//	@Failure		400		{object}	HTTPErrResp
//	@Failure		500		{object}	HTTPErrResp
//	@Router			/kc/login [post]
func (e *rest) KCLogin(c *gin.Context) {
	ctx := c.Request.Context()

	var requestBody UserLoginRequest

	if err := c.ShouldBind(&requestBody); err != nil {
		e.httpRespError(c, x.WrapWithCode(err, commonerr.CodeHTTPUnmarshal, "unmarshall_err"))
		return
	}

	if requestBody.Data.User == nil {
		e.httpRespError(c, x.NewWithCode(commonerr.CodeHTTPBadRequest, "invalid_payload"))
		return
	}

	data := requestBody.Data.User

	validationErr := data.Validate()
	if validationErr != nil {
		e.httpRespError(c, x.WrapWithCode(validationErr, commonerr.CodeHTTPBadRequest, "validation_error"))
		return
	}

	jwt, err := e.kc.Gocloak.Login(ctx, e.kc.ClientId, e.kc.ClientSecret, e.kc.Realm, data.Username, data.Password)
	if err != nil {
		e.httpRespError(c, x.NewWithCode(commonerr.CodeHTTPBadRequest, "invalid_payload"))
		return
	}

	result := dto.UserLoginDTO{
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		ExpiresIn:    jwt.ExpiresIn,
	}

	e.httpRespSuccess(c, http.StatusOK, result, nil)
}

// KC - GetEmployeeByID		godoc
//
//	@Summary		Get employee by ID
//	@Description	Endpoint for get a employee with ID
//	@Tags			KC
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			Authorization	header		string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Param			employeeID		path		string	true	"Employee ID"
//	@Success		200				{object}	HTTPEmployeeResp
//	@Failure		400				{object}	HTTPErrResp
//	@Failure		401				{object}	HTTPErrResp
//	@Failure		500				{object}	HTTPErrResp
//	@Router			/kc/employee/{employeeID} [get]
func (e *rest) KCGetEmployeeByID(c *gin.Context) {
	ctx := c.Request.Context()

	varID := c.Param("employeeID")

	employeeID, err := strconv.ParseInt(varID, 10, 64)
	if err != nil {
		e.httpRespError(c, x.Wrap(err, "decode_employee_id"))
		return
	}

	result, err := e.uc.Employee.KCGetEmployeeByID(ctx, employeeID)
	if err != nil {
		e.httpRespError(c, err)
		return
	}

	e.httpRespSuccess(c, http.StatusOK, result, nil)
}
