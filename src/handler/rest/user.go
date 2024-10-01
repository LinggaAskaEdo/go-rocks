package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/linggaaskaedo/go-rocks/src/business/dto"
	"github.com/linggaaskaedo/go-rocks/src/business/entity"
	"github.com/linggaaskaedo/go-rocks/src/common"
	commonerr "github.com/linggaaskaedo/go-rocks/stdlib/errors/common"
	x "github.com/linggaaskaedo/go-rocks/stdlib/errors/entity"
	"github.com/linggaaskaedo/go-rocks/stdlib/preference/sqlx"
)

// CreateUser		godoc
//
//	@Summary		Create user
//	@Description	Endpoint for creating a new user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			Authorization	header		string				true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Param			data			body		UserCreateRequest	true	"Payload create user"
//	@Success		201				{object}	HTTPUserResp
//	@Failure		400				{object}	HTTPErrResp
//	@Failure		401				{object}	HTTPErrResp
//	@Failure		429				{object}	HTTPErrResp
//	@Failure		500				{object}	HTTPErrResp
//	@Router			/user [post]
func (e *rest) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()

	var requestBody UserCreateRequest

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

	divisionID, err := common.MixerDecode(data.DivisionID)
	if err != nil {
		e.httpRespError(c, err)
		return
	}

	hashPassword, err := common.HashPassword(data.Password)
	if err != nil {
		e.httpRespError(c, x.WrapWithCode(err, commonerr.CodeHTTPBadRequest, "hash_password"))
		return
	}

	_, err = e.uc.Division.GetDivisioByID(ctx, divisionID)
	if err != nil {
		e.httpRespError(c, err)
		return
	}

	userEntity := entity.User{
		Username: data.Username,
		Email:    data.Email,
		Phone:    data.Phone,
		Division: entity.Division{
			ID: divisionID,
		},
		Password:  hashPassword,
		CreatedAt: sqlx.NullTime{Valid: true, Time: time.Now()},
	}

	result, err := e.uc.User.CreateUser(ctx, userEntity)
	if err != nil {
		e.httpRespError(c, err)
		return
	}

	e.httpRespSuccess(c, http.StatusCreated, result, nil)
}

// GetUserByID		godoc
//
//	@Summary		Get user by ID
//	@Description	Endpoint for get a user with ID
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			Authorization	header		string	true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Param			userID			path		string	true	"User ID"
//	@Success		200				{object}	HTTPUserResp
//	@Failure		400				{object}	HTTPErrResp
//	@Failure		401				{object}	HTTPErrResp
//	@Failure		500				{object}	HTTPErrResp
//	@Router			/user/{userID} [get]
func (e *rest) GetUserByID(c *gin.Context) {
	ctx := c.Request.Context()

	userPubID := c.Param("userID")

	userID, err := common.MixerDecode(userPubID)
	if err != nil {
		e.httpRespError(c, err)
		return
	}

	result, err := e.uc.User.GetUserByUserID(ctx, userID)
	if err != nil {
		e.httpRespError(c, err)
		return
	}

	e.httpRespSuccess(c, http.StatusOK, result, nil)
}

// UserLogin 	godoc
//
//	@Summary		Login user
//	@Description	Endpoint for generating access token for user
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			data	body		UserLoginRequest	true	"Payload user login"
//	@Success		200		{object}	HTTPUserLoginResp
//	@Failure		400		{object}	HTTPErrResp
//	@Failure		500		{object}	HTTPErrResp
//	@Router			/user/login [post]
func (e *rest) UserLogin(c *gin.Context) {
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

	resultUser, err := e.uc.User.GetUserByUsername(ctx, data.Username)
	if err != nil {
		e.httpRespError(c, err)
		return
	}

	if ok := common.CheckPasswordHash(data.Password, resultUser.HashPassword); !ok {
		e.httpRespError(c, x.WrapWithCode(x.New("wrong_password"), commonerr.CodeHTTPBadRequest, "check_password"))
		return
	}

	if err = e.auth.CheckToken(c, resultUser.PublicID); err != nil {
		e.httpRespError(c, err)
		return
	}

	generatedToken, err := e.auth.GenerateToken(c, resultUser)
	if err != nil {
		e.httpRespError(c, err)
		return
	}

	result := dto.UserLoginDTO{
		AccessToken:  generatedToken.AccessToken,
		RefreshToken: generatedToken.RefreshToken,
		ExpiresAt:    generatedToken.ExpiresAt,
		ExpiresRt:    generatedToken.ExpiresRt,
	}

	e.httpRespSuccess(c, http.StatusOK, result, nil)
}

// UserLogout		godoc
//
//	@Summary		User logout
//	@Description	Endpoint user for logging out
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			Authorization	header		string				true	"Insert your access token"	default(Bearer <Add access token here>)
//	@Param			data			body		UserLogoutRequest	true	"Payload user logout"
//	@Success		200				{object}	HTTPUserLogoutResp
//	@Failure		400				{object}	HTTPErrResp
//	@Failure		401				{object}	HTTPErrResp
//	@Failure		500				{object}	HTTPErrResp
//	@Router			/user/logout [post]
func (e *rest) UserLogout(c *gin.Context) {
	var requestBody UserLogoutRequest

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

	err := e.auth.InvalidateToken(c)
	if err != nil {
		e.httpRespError(c, x.WrapWithCode(err, commonerr.CodeHTTPBadRequest, "token_error"))
		return
	}

	result := dto.UserLogoutDTO{
		Message: "Succesfully logged out",
	}

	e.httpRespSuccess(c, http.StatusOK, result, nil)
}

// UserRelog		godoc
//
//	@Summary		User relog
//	@Description	Endpoint user for renewing access token
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Param			data	body		UserRelogRequest	true	"Payload user logout"
//	@Success		200		{object}	HTTPUserLoginResp
//	@Failure		400		{object}	HTTPErrResp
//	@Failure		401		{object}	HTTPErrResp
//	@Failure		500		{object}	HTTPErrResp
//	@Router			/user/relog [post]
func (e *rest) UserRelog(c *gin.Context) {
	var requestBody UserRelogRequest

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

	accessDetails, err := e.auth.ValidateRefreshToken(c, data.RefreshToken)
	if err != nil {
		e.httpRespError(c, x.WrapWithCode(err, commonerr.CodeHTTPBadRequest, "validate_token_error"))
		return
	}

	err = e.auth.InvalidateRefreshToken(c, accessDetails)
	if err != nil {
		e.httpRespError(c, x.WrapWithCode(err, commonerr.CodeHTTPBadRequest, "invalidate_token_error"))
		return
	}

	user := dto.UserDTO{
		PublicID: accessDetails.UserID,
		Username: accessDetails.Username,
	}

	generatedToken, err := e.auth.GenerateToken(c, user)
	if err != nil {
		e.httpRespError(c, err)
		return
	}

	result := dto.UserLoginDTO{
		AccessToken:  generatedToken.AccessToken,
		RefreshToken: generatedToken.RefreshToken,
		ExpiresAt:    generatedToken.ExpiresAt,
		ExpiresRt:    generatedToken.ExpiresRt,
	}

	e.httpRespSuccess(c, http.StatusOK, result, nil)
}
