package user

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"github.com/linggaaskaedo/go-rocks/src/business/entity"
)

type DomainItf interface {
	CreateUser(ctx context.Context, userEntity entity.User) (entity.User, error)
	GetUserByUserID(ctx context.Context, userID int64) (entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (entity.User, error)
}

type user struct {
	redis *redis.Client
	sql0  *sqlx.DB
}

type Options struct {
}

func InitUserDomain(redis *redis.Client, sql0 *sqlx.DB) DomainItf {
	return &user{
		redis: redis,
		sql0:  sql0,
	}
}
