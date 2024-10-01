package rest

import (
	"sync"

	"github.com/gin-gonic/gin"

	"github.com/linggaaskaedo/go-rocks/src/business/usecase"
	"github.com/linggaaskaedo/go-rocks/stdlib/auth"
	"github.com/linggaaskaedo/go-rocks/stdlib/keycloak"
	"github.com/linggaaskaedo/go-rocks/stdlib/middleware"
	libparser "github.com/linggaaskaedo/go-rocks/stdlib/parser"
)

var once = &sync.Once{}

type rest struct {
	gin   *gin.Engine
	auth  auth.Auth
	mw    middleware.Middleware
	param libparser.ParamParser
	kc    *keycloak.Keycloak
	uc    *usecase.Usecase
}

type Options struct {
}

func Init(gin *gin.Engine, auth auth.Auth, mw middleware.Middleware, parser libparser.Parser, kc *keycloak.Keycloak, uc *usecase.Usecase) {
	var e *rest

	once.Do(func() {
		e = &rest{
			gin:   gin,
			auth:  auth,
			mw:    mw,
			param: parser.ParamParser(),
			kc:    kc,
			uc:    uc,
		}

		e.Serve()
	})
}

func (e *rest) Serve() {
	// User
	e.gin.POST("/user", e.mw.JWT(), e.mw.Limiter("1-M", 3), e.CreateUser)
	e.gin.GET("/user/:userID", e.mw.JWT(), e.mw.Limiter("1-M", 3), e.GetUserByID)
	e.gin.POST("/user/login", e.mw.Limiter("1-M", 3), e.UserLogin)
	e.gin.POST("/user/logout", e.mw.JWT(), e.mw.Limiter("1-M", 3), e.UserLogout)
	e.gin.POST("/user/relog", e.mw.Limiter("1-M", 3), e.UserRelog)

	// Division
	e.gin.POST("/division", e.mw.JWT(), e.mw.Limiter("1-M", 3), e.CreateDivision)
	e.gin.GET("/division/:divisionID", e.mw.JWT(), e.mw.Limiter("1-M", 3), e.GetDivisionByID)
	e.gin.GET("/division", e.mw.JWT(), e.GetDivision)

	// Employee
	e.gin.GET("/employee/:employeeID", e.mw.JWT(), e.GetEmployeeByID)
	e.gin.GET("/employee", e.GetEmployee)

	// Keycloak
	e.gin.POST("/kc/login", e.KCLogin)
	e.gin.GET("/kc/employee/:employeeID", e.mw.KC(), e.KCGetEmployeeByID)
}
