package employee

import (
	"context"
	"database/sql"

	"github.com/rs/zerolog"

	"github.com/linggaaskaedo/go-rocks/src/business/entity"
	"github.com/linggaaskaedo/go-rocks/src/common"
	commonerr "github.com/linggaaskaedo/go-rocks/stdlib/errors/common"
	x "github.com/linggaaskaedo/go-rocks/stdlib/errors/entity"
)

func (e *employee) getSQLEmployeeByID(ctx context.Context, employeeID int64) (entity.Employee, error) {
	var result entity.Employee

	query, args := GetEmployeeByID(employeeID)
	row := e.sql0.QueryRowContext(ctx, query, args...)

	if err := row.Scan(
		&result.EmpNo,
		&result.FirstName,
		&result.LastName,
		&result.Gender,
		&result.HireDate,
	); err != nil {
		if err == sql.ErrNoRows {
			return result, x.WrapWithCode(err, commonerr.CodeSQLEmptyRow, "get_employee_by_id")
		}

		return result, x.WrapWithCode(err, commonerr.CodeSQLRowScan, "get_employee_by_id")
	}

	zerolog.Ctx(ctx).Debug().Any("result", result).Send()

	return result, nil
}

func (e *employee) getSQLEmployee(ctx context.Context, param entity.EmployeeParam) ([]entity.Employee, entity.Pagination, error) {
	var (
		results      = make([]entity.Employee, 0)
		totalRecords int64
	)

	param.Limit = common.ValidateLimit(param.Limit)
	param.Page = common.ValidatePage(param.Page)

	pagination := entity.Pagination{
		CurrentPage:     param.Page,
		CurrentElements: 0,
		TotalPages:      0,
		TotalElements:   0,
		SortBy:          []string{},
	}

	query, queryArgs, queryCount, queryCountArgs := GetEmployee(param)

	rows, err := e.sql0.QueryContext(ctx, query, queryArgs...)
	if err != nil {
		return results, pagination, x.WrapWithCode(err, commonerr.CodeSQLRead, "get_employee")
	}
	defer rows.Close()

	for rows.Next() {
		result := entity.Employee{}

		if err := rows.Scan(
			&result.EmpNo,
			&result.FirstName,
			&result.LastName,
			&result.Gender,
		); err != nil {
			return results, pagination, x.WrapWithCode(err, commonerr.CodeSQLRowScan, "get_employee")
		}

		results = append(results, result)
	}

	err = e.sql0.GetContext(ctx, &totalRecords, queryCount, queryCountArgs...)
	if err != nil {
		return results, pagination, x.WrapWithCode(err, commonerr.CodeSQLRead, "get_employee")
	}

	zerolog.Ctx(ctx).Debug().Any("totalRecords", totalRecords).Send()

	// Update Pagination
	totalPage := totalRecords / param.Limit
	if totalRecords%param.Limit > 0 || totalRecords == 0 {
		totalPage++
	}

	pagination.TotalPages = common.ValidatePage(totalPage)
	pagination.CurrentElements = int64(len(results))
	pagination.TotalElements = totalRecords
	pagination.SortBy = param.SortBy

	return results, pagination, nil
}
