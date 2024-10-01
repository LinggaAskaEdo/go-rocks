package usecase

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"github.com/linggaaskaedo/go-rocks/src/business/domain"
	"github.com/linggaaskaedo/go-rocks/src/business/usecase/division"
	"github.com/linggaaskaedo/go-rocks/src/business/usecase/employee"
	"github.com/linggaaskaedo/go-rocks/src/business/usecase/user"
)

type Usecase struct {
	User     user.UsecaseItf
	Division division.UsecaseItf
	Employee employee.UsecaseItf
}

type Options struct {
}

func Init(
	redis *redis.Client,
	sql0 *sqlx.DB,
	dom *domain.Domain,
) *Usecase {
	return &Usecase{
		User: user.InitUserUsecase(
			redis,
			sql0,
			dom.User,
		),
		Division: division.InitDivisionUsecase(
			redis,
			sql0,
			dom.Division,
		),
		Employee: employee.InitEmployeeUsecase(
			redis,
			sql0,
			dom.Employee,
		),
	}
}
