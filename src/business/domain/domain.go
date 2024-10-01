package domain

import (
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"

	"github.com/linggaaskaedo/go-rocks/src/business/domain/division"
	"github.com/linggaaskaedo/go-rocks/src/business/domain/employee"
	"github.com/linggaaskaedo/go-rocks/src/business/domain/user"
)

type Domain struct {
	User     user.DomainItf
	Division division.DomainItf
	Employee employee.DomainItf
}

func Init(
	redis *redis.Client,
	sql0 *sqlx.DB,
) *Domain {
	return &Domain{
		User: user.InitUserDomain(
			redis,
			sql0,
		),
		Division: division.InitDivisionDomain(
			redis,
			sql0,
		),
		Employee: employee.InitEmployeeDomain(
			redis,
			sql0,
		),
	}
}
