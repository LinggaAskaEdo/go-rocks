package user

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	usr "github.com/linggaaskaedo/go-rocks/src/business/domain/user"
	"github.com/linggaaskaedo/go-rocks/src/business/dto"
	"github.com/linggaaskaedo/go-rocks/src/business/entity"
)

type UsecaseItf interface {
	CreateUser(ctx context.Context, userEntity entity.User) (dto.UserDTO, error)
	GetUserByUserID(ctx context.Context, userID int64) (dto.UserDTO, error)
	GetUserByUsername(ctx context.Context, username string) (dto.UserDTO, error)
}

type user struct {
	redis *redis.Client
	sql0  *sqlx.DB
	user  usr.DomainItf
}

type Options struct {
}

func InitUserUsecase(redis *redis.Client, sql0 *sqlx.DB, u usr.DomainItf) UsecaseItf {
	return &user{
		redis: redis,
		sql0:  sql0,
		user:  u,
	}
}
