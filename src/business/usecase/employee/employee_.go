package employee

import (
	"context"
	"slices"
	"sort"

	"github.com/linggaaskaedo/go-rocks/src/business/dto"
	"github.com/linggaaskaedo/go-rocks/src/business/entity"
	"github.com/linggaaskaedo/go-rocks/src/common"
)

func (e *employee) GetEmployeeByID(ctx context.Context, employeeID int64) (dto.EmployeeDTO, error) {
	var employeeDTO dto.EmployeeDTO

	result, err := e.employee.GetEmployeeByID(ctx, employeeID)
	if err != nil {
		return employeeDTO, err
	}

	employeeDTO.PublicID = common.MixerEncode(result.EmpNo)
	employeeDTO.FirstName = result.FirstName
	employeeDTO.LastName = result.LastName
	employeeDTO.Gender = result.Gender

	return employeeDTO, nil
}

func (e *employee) GetEmployee(ctx context.Context, cacheControl entity.CacheControl, param entity.EmployeeParam) ([]dto.EmployeeDTO, entity.Pagination, error) {
	var results = make([]dto.EmployeeDTO, 0)
	var pagination entity.Pagination

	employees, pagination, err := e.employee.GetEmployee(ctx, cacheControl, param)
	if err != nil {
		return results, pagination, err
	}

	for _, employee := range employees {
		data := dto.EmployeeDTO{
			PublicID:  common.MixerEncode(employee.EmpNo),
			FirstName: employee.FirstName,
			LastName:  employee.LastName,
			Gender:    employee.Gender,
		}

		results = append(results, data)
	}

	// since id already masking, sort slice by id do here
	if slices.Contains(param.SortBy, "id") {
		sort.Slice(results, func(i, j int) bool {
			if results[i].PublicID != results[j].PublicID {
				return results[i].PublicID < results[j].PublicID
			}

			return results[i].PublicID < results[j].PublicID
		})
	}

	if slices.Contains(param.SortBy, "-id") {
		sort.Slice(results, func(i, j int) bool {
			if results[i].PublicID != results[j].PublicID {
				return results[i].PublicID > results[j].PublicID
			}

			return results[i].PublicID > results[j].PublicID
		})
	}

	return results, pagination, nil
}

func (e *employee) KCGetEmployeeByID(ctx context.Context, employeeID int64) (dto.KCEmployeeDTO, error) {
	var employeeDTO dto.KCEmployeeDTO

	result, err := e.employee.GetEmployeeByID(ctx, employeeID)
	if err != nil {
		return employeeDTO, err
	}

	employeeDTO.ID = result.EmpNo
	employeeDTO.FirstName = result.FirstName
	employeeDTO.LastName = result.LastName
	employeeDTO.Gender = result.Gender
	employeeDTO.HireDate = common.StringTime(result.HireDate)

	return employeeDTO, nil
}
