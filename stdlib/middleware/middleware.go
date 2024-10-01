package middleware

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"

	"github.com/linggaaskaedo/go-rocks/stdlib/auth"
	"github.com/linggaaskaedo/go-rocks/stdlib/keycloak"
	"github.com/linggaaskaedo/go-rocks/stdlib/preference"
)

const TimeFormat = "2006-01-02 15:04:05"

var (
	once = &sync.Once{}

	timeDict = map[string]time.Duration{
		"S": time.Second,
		"M": time.Minute,
		"H": time.Hour,
		"D": time.Hour * 24,
	}
)

type Middleware interface {
	Handler() gin.HandlerFunc
	CORS() gin.HandlerFunc
	Limiter(command string, limit int) gin.HandlerFunc
	JWT() gin.HandlerFunc
	KC() gin.HandlerFunc
}

type middleware struct {
	log       zerolog.Logger
	auth      auth.Auth
	kc        *keycloak.Keycloak
	limit     int
	deadline  int64
	shaScript map[string]string
	period    time.Duration
	rdb       *redis.Client
}

type Options struct {
	Limiter LimiterOptions
}

type LimiterOptions struct {
	Command string
	Limit   int
}

func Init(log zerolog.Logger, opt Options, auth auth.Auth, kc *keycloak.Keycloak, rdb *redis.Client) Middleware {
	var m *middleware

	once.Do(func() {
		var limit int
		var period time.Duration
		var shaScript map[string]string

		values := strings.Split(opt.Limiter.Command, "-")
		if len(values) != 2 {
			log.Panic().Err(errors.New(preference.FormatError)).Send()
		}

		unit, err := strconv.Atoi(values[0])
		if err != nil {
			log.Panic().Err(errors.New(preference.FormatError)).Send()
		}

		if unit <= 0 {
			log.Panic().Err(errors.New(preference.CommandError)).Send()
		}

		if t, ok := timeDict[strings.ToUpper(values[1])]; ok {
			period = time.Duration(unit) * t
		} else {
			log.Panic().Err(errors.New(preference.FormatError)).Send()
		}

		if opt.Limiter.Limit <= 0 {
			log.Panic().Err(errors.New(preference.LimitError)).Send()
		}

		limit = opt.Limiter.Limit

		resetSHA, err := rdb.ScriptLoad(context.Background(), ResetScript).Result()
		if err != nil {
			log.Panic().Err(err).Send()
		}

		normalSHA, err := rdb.ScriptLoad(context.Background(), Script).Result()
		if err != nil {
			log.Panic().Err(err).Send()
		}

		shaScript = make(map[string]string)
		shaScript["reset"] = resetSHA
		shaScript["normal"] = normalSHA

		m = &middleware{
			log:       log,
			auth:      auth,
			kc:        kc,
			limit:     limit,
			shaScript: shaScript,
			period:    period,
			rdb:       rdb,
		}
	})

	return m
}
