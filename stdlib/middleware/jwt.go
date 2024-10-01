package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	commonerr "github.com/linggaaskaedo/go-rocks/stdlib/errors/common"
	x "github.com/linggaaskaedo/go-rocks/stdlib/errors/entity"
)

func (mw *middleware) JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := mw.auth.ValidateToken(c)
		if err != nil {
			errorResp := mw.httpRespError(c, x.WrapWithCode(err, commonerr.CodeHTTPUnauthorized, "This user does not have account info"))
			c.JSON(http.StatusUnauthorized, errorResp)
			c.Abort()
			return
		}

		if token == nil {
			errorResp := mw.httpRespError(c, x.NewWithCode(commonerr.CodeHTTPUnauthorized, "Invalid token"))
			c.JSON(http.StatusUnauthorized, errorResp)
			c.Abort()
			return
		}

		c.Header("userID", token.UserID)
		c.Header("username", token.Username)
		c.Next()
	}
}
