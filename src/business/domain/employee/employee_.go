package employee

import (
	"context"

	"github.com/linggaaskaedo/go-rocks/src/business/entity"
)

func (e *employee) GetEmployeeByID(ctx context.Context, employeeID int64) (entity.Employee, error) {
	result, err := e.getSQLEmployeeByID(ctx, employeeID)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (e *employee) GetEmployee(ctx context.Context, cacheControl entity.CacheControl, param entity.EmployeeParam) ([]entity.Employee, entity.Pagination, error) {
	result, pagination, err := e.getSQLEmployee(ctx, param)
	if err != nil {
		return result, pagination, err
	}

	return result, pagination, nil
}
