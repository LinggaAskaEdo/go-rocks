package division

import (
	"context"

	"github.com/linggaaskaedo/go-rocks/src/business/dto"
	"github.com/linggaaskaedo/go-rocks/src/business/entity"
	"github.com/linggaaskaedo/go-rocks/src/common"
)

func (d *division) CreateDivision(ctx context.Context, divisionEntity entity.Division) (dto.DivisionDTO, error) {
	var divisionDTO dto.DivisionDTO

	result, err := d.division.CreateDivision(ctx, divisionEntity)
	if err != nil {
		return divisionDTO, err
	}

	divisionDTO.PublicID = common.MixerEncode(result.ID)
	divisionDTO.Name = result.Name
	divisionDTO.IsDeleted = result.IsDeleted

	return divisionDTO, nil
}

func (d *division) GetDivisioByID(ctx context.Context, divisionID int64) (dto.DivisionDTO, error) {
	var divisionDTO dto.DivisionDTO

	result, err := d.division.GetDivisioByID(ctx, divisionID)
	if err != nil {
		return divisionDTO, err
	}

	divisionDTO.PublicID = common.MixerEncode(result.ID)
	divisionDTO.Name = result.Name
	divisionDTO.IsDeleted = result.IsDeleted

	return divisionDTO, nil
}

func (d *division) GetDivision(ctx context.Context, cacheControl entity.CacheControl, param entity.DivisionParam) ([]dto.DivisionDTO, entity.Pagination, error) {
	var results = make([]dto.DivisionDTO, 0)
	var pagination entity.Pagination

	divisions, pagination, err := d.division.GetDivision(ctx, cacheControl, param)
	if err != nil {
		return results, pagination, err
	}

	for _, division := range divisions {
		data := dto.DivisionDTO{
			PublicID:  common.MixerEncode(division.ID),
			Name:      division.Name,
			IsDeleted: division.IsDeleted,
		}

		results = append(results, data)
	}

	return results, pagination, nil
}
