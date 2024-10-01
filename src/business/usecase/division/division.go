package division

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	div "github.com/linggaaskaedo/go-rocks/src/business/domain/division"
	"github.com/linggaaskaedo/go-rocks/src/business/dto"
	"github.com/linggaaskaedo/go-rocks/src/business/entity"
)

type UsecaseItf interface {
	CreateDivision(ctx context.Context, divisionEntity entity.Division) (dto.DivisionDTO, error)
	GetDivisioByID(ctx context.Context, divisionID int64) (dto.DivisionDTO, error)
	GetDivision(ctx context.Context, cacheControl entity.CacheControl, param entity.DivisionParam) ([]dto.DivisionDTO, entity.Pagination, error)
}

type division struct {
	redis    *redis.Client
	sql0     *sqlx.DB
	division div.DomainItf
}

type Options struct {
}

func InitDivisionUsecase(redis *redis.Client, sql0 *sqlx.DB, d div.DomainItf) UsecaseItf {
	return &division{
		redis:    redis,
		sql0:     sql0,
		division: d,
	}
}
