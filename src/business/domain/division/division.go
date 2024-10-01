package division

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"github.com/linggaaskaedo/go-rocks/src/business/entity"
)

type DomainItf interface {
	CreateDivision(ctx context.Context, divisionEntity entity.Division) (entity.Division, error)
	GetDivisioByID(ctx context.Context, divisionID int64) (entity.Division, error)
	GetDivision(ctx context.Context, cacheControl entity.CacheControl, param entity.DivisionParam) ([]entity.Division, entity.Pagination, error)
}

type division struct {
	redis *redis.Client
	sql0  *sqlx.DB
}

type Options struct {
}

func InitDivisionDomain(redis *redis.Client, sql0 *sqlx.DB) DomainItf {
	return &division{
		redis: redis,
		sql0:  sql0,
	}
}
