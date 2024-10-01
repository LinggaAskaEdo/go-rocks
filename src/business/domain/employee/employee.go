package employee

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"github.com/linggaaskaedo/go-rocks/src/business/entity"
)

type DomainItf interface {
	GetEmployeeByID(ctx context.Context, employeeID int64) (entity.Employee, error)
	GetEmployee(ctx context.Context, cacheControl entity.CacheControl, param entity.EmployeeParam) ([]entity.Employee, entity.Pagination, error)
}

type employee struct {
	redis *redis.Client
	sql0  *sqlx.DB
}

type Options struct {
}

func InitEmployeeDomain(redis *redis.Client, sql0 *sqlx.DB) DomainItf {
	return &employee{
		redis: redis,
		sql0:  sql0,
	}
}
