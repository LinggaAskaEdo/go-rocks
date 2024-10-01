package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/linggaaskaedo/go-rocks/docs"
	"github.com/linggaaskaedo/go-rocks/stdlib/middleware"
)

type Options struct {
}

func Init(log zerolog.Logger, opt Options, middleware middleware.Middleware) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(middleware.Handler())
	router.Use(middleware.CORS())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.DefaultModelsExpandDepth(-1)))

	return router
}
