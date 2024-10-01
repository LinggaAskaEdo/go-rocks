package division

import (
	"context"
	"database/sql"
	"regexp"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"

	"github.com/linggaaskaedo/go-rocks/src/business/entity"
	"github.com/linggaaskaedo/go-rocks/src/common"
	commonerr "github.com/linggaaskaedo/go-rocks/stdlib/errors/common"
	x "github.com/linggaaskaedo/go-rocks/stdlib/errors/entity"
)

func (d *division) createSQLDivision(ctx context.Context, tx *sqlx.Tx, divisionEntity entity.Division) (*sqlx.Tx, entity.Division, error) {
	row, err := tx.Exec(CreateDivision, divisionEntity.Name, divisionEntity.CreatedAt)
	if err != nil {
		return tx, divisionEntity, x.Wrap(err, "create_division")
	}

	divisionEntity.ID, err = row.LastInsertId()
	zerolog.Ctx(ctx).Debug().Any("result", divisionEntity).Send()

	return tx, divisionEntity, nil
}

func (d *division) getSQLDivisionByID(ctx context.Context, divisionID int64) (entity.Division, error) {
	var result entity.Division

	row := d.sql0.QueryRowContext(ctx, GetDivisionByID, divisionID)

	if err := row.Scan(
		&result.ID,
		&result.Name,
		&result.IsDeleted,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return result, x.WrapWithCode(err, commonerr.CodeSQLEmptyRow, "get_division_by_id")
		}

		return result, x.WrapWithCode(err, commonerr.CodeSQLRowScan, "get_division_by_id")
	}

	zerolog.Ctx(ctx).Debug().Any("result", result).Send()

	return result, nil
}

func (d *division) getSQLDivision(ctx context.Context, param entity.DivisionParam) ([]entity.Division, entity.Pagination, error) {
	var (
		results      = make([]entity.Division, 0)
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

	qb := common.NewSQLClauseBuilder("param", "db", "", param.Page, param.Limit).AliasPrefix("division", &param)

	queryExt, sortBy, args, err := qb.Build()
	if err != nil {
		return results, pagination, x.WrapWithCode(err, commonerr.CodeSQLBuilder, "get_division")
	}

	queryExtC, argsCount := queryExt, args

	query, args, err := sqlx.In(GetDivision+queryExt, args...)
	if err != nil {
		return results, pagination, x.WrapWithCode(err, commonerr.CodeSQLBuilder, "get_division")
	}

	zerolog.Ctx(ctx).Debug().Any("QUERY", query).Send()

	rows, err := d.sql0.QueryContext(ctx, query, args...)
	if err != nil {
		return results, pagination, x.WrapWithCode(err, commonerr.CodeSQLRead, "get_division")
	}
	defer rows.Close()

	for rows.Next() {
		result := entity.Division{}

		if err := rows.Scan(
			&result.ID,
			&result.Name,
			&result.IsDeleted,
			&result.CreatedAt,
			&result.UpdatedAt,
			&result.DeletedAt,
		); err != nil {
			return results, pagination, x.WrapWithCode(err, commonerr.CodeSQLRowScan, "get_division")
		}

		results = append(results, result)
	}

	queryCount, argsCount, err := sqlx.In(CountDivision+queryExtC, argsCount...)
	if err != nil {
		return results, pagination, x.WrapWithCode(err, commonerr.CodeSQLBuilder, "get_division")
	}

	reg := regexp.MustCompile(`(.*)LIMIT.*;`)
	queryCount = reg.ReplaceAllString(queryCount, "${1};")

	err = d.sql0.GetContext(ctx, &totalRecords, queryCount, argsCount...)
	if err != nil {
		return results, pagination, x.WrapWithCode(err, commonerr.CodeSQLRead, "get_division")
	}

	// Update Pagination
	totalPage := totalRecords / param.Limit
	if totalRecords%param.Limit > 0 || totalRecords == 0 {
		totalPage++
	}

	pagination.TotalPages = common.ValidatePage(totalPage)
	pagination.CurrentElements = int64(len(results))
	pagination.TotalElements = totalRecords
	pagination.SortBy = sortBy

	return results, pagination, nil
}
