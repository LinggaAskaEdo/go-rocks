package auth

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/segmentio/ksuid"

	commonerr "github.com/linggaaskaedo/go-rocks/stdlib/errors/common"
	x "github.com/linggaaskaedo/go-rocks/stdlib/errors/entity"
)

type TokenDetails struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	AccessUUID   string `json:"-"`
	RefreshUUID  string `json:"-"`
	ExpiresAt    int64  `json:"expiresAt"`
	ExpiresRt    int64  `json:"expiresRt"`
}

type AccessDetails struct {
	AccessUUID  string
	RefreshUUID string
	UserID      string
	Username    string
}

func (a *auth) GenerateToken(c *gin.Context, data interface{}) (*TokenDetails, error) {
	ctx := c.Request.Context()

	td := &TokenDetails{}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(a.privateKey)
	if err != nil {
		return nil, x.Wrap(err, "parse_key")
	}

	dataVal := reflect.ValueOf(data)
	publicID := dataVal.FieldByName("PublicID").String()
	username := dataVal.FieldByName("Username").String()

	td.ExpiresAt = time.Now().Add(a.expiredToken).Unix()
	td.AccessUUID = ksuid.New().String()

	td.ExpiresRt = time.Now().Add(a.expiredRefreshToken).Unix()
	td.RefreshUUID = td.AccessUUID + "++" + publicID

	at := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"exp":         td.ExpiresAt,
		"access_uuid": td.AccessUUID,
		"user_id":     publicID,
		"name":        username,
		"authorized":  true,
	})

	td.AccessToken, err = at.SignedString(key)
	if err != nil {
		return nil, x.Wrap(err, "sign_access_token")
	}

	rt := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"exp":          td.ExpiresRt,
		"refresh_uuid": td.RefreshUUID,
		"user_id":      publicID,
		"name":         username,
	})

	td.RefreshToken, err = rt.SignedString(key)
	if err != nil {
		return nil, x.Wrap(err, "sign_refresh_token")
	}

	err = a.saveToRedis(ctx, publicID, td)
	if err != nil {
		return nil, err
	}

	return td, nil
}

func (a *auth) ValidateToken(c *gin.Context) (*AccessDetails, error) {
	return a.checkingToken(c)
}

func (a *auth) ValidateRefreshToken(c *gin.Context, tokenStr string) (*AccessDetails, error) {
	ctx := c.Request.Context()

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

	var accessUUID, refreshUUID, redisIDUser string

	refreshUUID, ok = claims["refresh_uuid"].(string)
	if !ok {
		return nil, err
	}

	redisIDUser, err = a.redis.Get(ctx, refreshUUID).Result()
	if err != nil {
		return nil, x.New("redis_get_token")
	}

	if userID != redisIDUser {
		return nil, x.New("Authentification failure")
	}

	return &AccessDetails{
		AccessUUID:  accessUUID,
		RefreshUUID: refreshUUID,
		UserID:      redisIDUser,
		Username:    username,
	}, nil
}

func (a *auth) InvalidateToken(c *gin.Context) error {
	ctx := c.Request.Context()

	accessDetails, err := a.checkingToken(c)
	if err != nil {
		return x.Wrap(err, "store_redis_access_token")
	}

	return a.deleteToRedis(ctx, accessDetails, false)
}

func (a *auth) InvalidateRefreshToken(c *gin.Context, accessDetails *AccessDetails) error {
	ctx := c.Request.Context()

	return a.deleteToRedis(ctx, accessDetails, true)
}

func (a *auth) CheckToken(c *gin.Context, userID string) error {
	ctx := c.Request.Context()

	var cursorRefToken, cursorToken uint64

	cursorRefToken = 0
	cursorToken = 0

	for {
		arrRefreshToken, cursorRefToken, err := a.redis.Scan(ctx, cursorRefToken, fmt.Sprintf("*%s", userID), 1000).Result()
		if err != nil {
			return err
		}

		// Found refresh token in redis
		if cursorRefToken == 0 {
			if len(arrRefreshToken) >= 1 {
				// Checking possibilty token still exist in redis
				refershToken := arrRefreshToken[0]
				token := strings.ReplaceAll(refershToken, "++"+userID, "")

				arrToken, cursorToken, err := a.redis.Scan(ctx, cursorToken, token, 1000).Result()
				if err != nil {
					return err
				}

				if cursorToken == 0 {
					if len(arrToken) >= 1 {
						return x.NewWithCode(commonerr.CodeTokenStillValid, "Token still valid")
					} else {
						return x.NewWithCode(commonerr.CodeTokenRefreshStillValid, "Refresh token still valid, please re-log")
					}
				} else {
					continue
				}
			} else {
				return nil
			}
		} else {
			continue
		}
	}
}
