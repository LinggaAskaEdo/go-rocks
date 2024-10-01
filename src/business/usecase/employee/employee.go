package employee

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	emp "github.com/linggaaskaedo/go-rocks/src/business/domain/employee"
	"github.com/linggaaskaedo/go-rocks/src/business/dto"
	"github.com/linggaaskaedo/go-rocks/src/business/entity"
)

type UsecaseItf interface {
	GetEmployeeByID(ctx context.Context, employeeID int64) (dto.EmployeeDTO, error)
	GetEmployee(ctx context.Context, cacheControl entity.CacheControl, param entity.EmployeeParam) ([]dto.EmployeeDTO, entity.Pagination, error)

	// KC
	KCGetEmployeeByID(ctx context.Context, employeeID int64) (dto.KCEmployeeDTO, error)
}

type employee struct {
	redis    *redis.Client
	sql0     *sqlx.DB
	employee emp.DomainItf
}

type Options struct {
}

func InitEmployeeUsecase(redis *redis.Client, sql0 *sqlx.DB, d emp.DomainItf) UsecaseItf {
	return &employee{
		redis:    redis,
		sql0:     sql0,
		employee: d,
	}
}
