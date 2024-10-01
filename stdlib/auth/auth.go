package auth

import (
	"os"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
)

var once = &sync.Once{}

type Auth interface {
	GenerateToken(c *gin.Context, data interface{}) (*TokenDetails, error)
	ValidateToken(c *gin.Context) (*AccessDetails, error)
	ValidateRefreshToken(c *gin.Context, token string) (*AccessDetails, error)
	InvalidateToken(c *gin.Context) error
	InvalidateRefreshToken(c *gin.Context, accessDetails *AccessDetails) error
	CheckToken(c *gin.Context, userID string) error
}

type auth struct {
	log                 zerolog.Logger
	redis               *redis.Client
	privateKey          []byte
	publicKey           []byte
	expiredToken        time.Duration
	expiredRefreshToken time.Duration
}

type Options struct {
	PrivateKey          string
	PublicKey           string
	ExpiredToken        int
	ExpiredRefreshToken int
	AutoLogoff          int
}

func Init(log zerolog.Logger, opt Options, redis *redis.Client) Auth {
	var a *auth

	once.Do(func() {
		privateKey, err := os.ReadFile(opt.PrivateKey)
		if err != nil {
			log.Panic().Err(err).Send()
		}

		publicKey, err := os.ReadFile(opt.PublicKey)
		if err != nil {
			log.Panic().Err(err).Send()
		}

		a = &auth{
			log:                 log,
			redis:               redis,
			privateKey:          privateKey,
			publicKey:           publicKey,
			expiredToken:        time.Duration(opt.ExpiredToken * int(time.Minute)),
			expiredRefreshToken: time.Duration(opt.ExpiredRefreshToken * int(time.Minute)),
		}
	})

	return a
}
