package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"

	commonerr "github.com/linggaaskaedo/go-rocks/stdlib/errors/common"
	x "github.com/linggaaskaedo/go-rocks/stdlib/errors/entity"
)

func (mw *middleware) extractBearerToken(token string) string {
	return strings.Replace(token, "Bearer ", "", 1)
}

func (mw *middleware) KC() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		auth := c.Request.Header["Authorization"]

		if len(auth) < 1 {
			errorResp := mw.httpRespError(c, x.NewWithCode(commonerr.CodeHTTPUnauthorized, "Authorization header missing"))
			c.JSON(http.StatusUnauthorized, errorResp)
			c.Abort()
			return
		}

		token := auth[0]

		if token == "" {
			errorResp := mw.httpRespError(c, x.NewWithCode(commonerr.CodeHTTPUnauthorized, "Authorization header missing"))
			c.JSON(http.StatusUnauthorized, errorResp)
			c.Abort()
			return
		}

		// extract Bearer token
		token = mw.extractBearerToken(token)

		result, err := mw.kc.Gocloak.RetrospectToken(ctx, token, mw.kc.ClientId, mw.kc.ClientSecret, mw.kc.Realm)
		if err != nil {
			errorResp := mw.httpRespError(c, x.NewWithCode(commonerr.CodeHTTPUnauthorized, "Invalid token"))
			c.JSON(http.StatusUnauthorized, errorResp)
			c.Abort()
			return
		}

		mw.log.Debug().Any("RESULT", result).Send()

		// check if the token isn't expired and valid
		if !*result.Active {
			errorResp := mw.httpRespError(c, x.NewWithCode(commonerr.CodeHTTPUnauthorized, "Invalid or expired Token"))
			c.JSON(http.StatusUnauthorized, errorResp)
			c.Abort()
			return
		}

		jwt, _, err := mw.kc.Gocloak.DecodeAccessToken(ctx, token, mw.kc.Realm)
		if err != nil {
			errorResp := mw.httpRespError(c, x.NewWithCode(commonerr.CodeHTTPUnauthorized, "Invalid or malformed token"))
			c.JSON(http.StatusUnauthorized, errorResp)
			c.Abort()
			return
		}

		jwtj, _ := json.Marshal(jwt)
		mw.log.Debug().Any("TOKEN", string(jwtj)).Send()

		uf, err := mw.kc.Gocloak.GetUserInfo(ctx, token, mw.kc.Realm)
		if err != nil {
			errorResp := mw.httpRespError(c, x.NewWithCode(commonerr.CodeHTTPUnauthorized, "Invalid or malformed token"))
			c.JSON(http.StatusUnauthorized, errorResp)
			c.Abort()
			return
		}

		mw.log.Debug().Any("USER_INFO", uf.String()).Send()

		groups, err := mw.kc.Gocloak.GetUserGroups(ctx, token, mw.kc.Realm, *uf.Sub, gocloak.GetGroupsParams{})
		if err != nil {
			mw.log.Debug().Any("ERROR", err).Send()
			errorResp := mw.httpRespError(c, x.NewWithCode(commonerr.CodeHTTPUnauthorized, "Invalid or malformed token"))
			c.JSON(http.StatusUnauthorized, errorResp)
			c.Abort()
			return
		}

		mw.log.Debug().Any("USER_GROUP", groups).Send()
	}
}
