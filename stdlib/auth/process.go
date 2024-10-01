package auth

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"

	commonerr "github.com/linggaaskaedo/go-rocks/stdlib/errors/common"
	x "github.com/linggaaskaedo/go-rocks/stdlib/errors/entity"
)

func (a *auth) checkingToken(c *gin.Context) (*AccessDetails, error) {
	ctx := c.Request.Context()

	tokenStr := a.extractToken(c)
	token, err := a.verifyToken(tokenStr)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, x.New("invalid_token")
	}

	userID := claims["user_id"].(string)
	username := claims["name"].(string)

	var accessUUID, redisIDUser string

	accessUUID, ok = claims["access_uuid"].(string)
	if !ok {
		return nil, x.WrapWithCode(err, commonerr.CodeHTTPUnauthorized, "failed claims accessUUID")
	}

	redisIDUser, err = a.redis.Get(ctx, accessUUID).Result()
	if err != nil {
		return nil, x.New("redis_get_token")
	}

	if userID != redisIDUser {
		return nil, x.New("Authentification failure")
	}

	return &AccessDetails{
		AccessUUID: accessUUID,
		UserID:     redisIDUser,
		Username:   username,
	}, nil
}

func (a *auth) extractToken(c *gin.Context) string {
	bearToken := c.Request.Header["Authorization"][0]
	if len(bearToken) == 0 {
		return ""
	}

	tokenArr := strings.Split(bearToken, " ")
	if len(tokenArr) == 2 {
		return tokenArr[1]
	}

	return ""
}

func (a *auth) verifyToken(tokenStr string) (*jwt.Token, error) {
	key, err := jwt.ParseRSAPublicKeyFromPEM(a.publicKey)
	if err != nil {
		return nil, x.Wrap(err, "parse_key")
	}

	token, err := jwt.Parse(tokenStr, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			s := fmt.Sprintf("unexpected signing method: %v", jwtToken.Header["alg"])
			return nil, x.Wrap(errors.New(s), "verify_token")
		}

		return key, nil
	})
	if err != nil {
		return nil, x.Wrap(err, "parse_token")
	}

	return token, nil
}
