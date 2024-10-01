package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"

	"github.com/linggaaskaedo/go-rocks/stdlib/preference"
)

type Options struct {
	Enabled         bool
	Network         string
	Address         string
	Password        string
	MaxRetries      int
	MinRetryBackoff time.Duration
	MaxRetryBackoff time.Duration
	DialTimeout     time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	PoolSize        int
	MinIdleConns    int
	MaxIdleConns    int
	MaxActiveConns  int
	PoolTimeout     time.Duration
}

func Init(log zerolog.Logger, opt Options, redisType string) *redis.Client {
	var redisClient *redis.Client
	var DB int

	if !opt.Enabled {
		return nil
	}

	if redisType == preference.REDIS_APPS {
		DB = 0
	} else if redisType == preference.REDIS_AUTH {
		DB = 11
	} else {
		DB = 13
	}

	redisClient = redis.NewClient(&redis.Options{
		Network:         opt.Network,
		Addr:            opt.Address,
		Password:        opt.Password,
		DB:              DB,
		MaxRetries:      opt.MaxRetries,
		MinRetryBackoff: opt.MinRetryBackoff,
		MaxRetryBackoff: opt.MaxRetryBackoff,
		DialTimeout:     opt.DialTimeout,
		ReadTimeout:     opt.ReadTimeout,
		WriteTimeout:    opt.WriteTimeout,
		PoolSize:        opt.PoolSize,
		MinIdleConns:    opt.MinIdleConns,
		MaxIdleConns:    opt.MaxIdleConns,
		MaxActiveConns:  opt.MaxActiveConns,
		PoolTimeout:     opt.PoolTimeout,
	})

	ping, err := redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Panic().Err(err).Msg(fmt.Sprintf("REDIS %s status: %s", redisType, ping))
	}

	log.Debug().Msg(fmt.Sprintf("REDIS %s status: %s", redisType, ping))

	return redisClient
}
