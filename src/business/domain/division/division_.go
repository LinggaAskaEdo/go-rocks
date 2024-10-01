package division

import (
	"context"
	"database/sql"

	"github.com/redis/go-redis/v9"

	"github.com/linggaaskaedo/go-rocks/src/business/entity"
	x "github.com/linggaaskaedo/go-rocks/stdlib/errors/entity"
	"github.com/rs/zerolog"
)

func (d *division) CreateDivision(ctx context.Context, divisionEntity entity.Division) (entity.Division, error) {
	tx, err := d.sql0.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return divisionEntity, x.Wrap(err, "tx_create_division")
	}

	tx, divisionEntity, err = d.createSQLDivision(ctx, tx, divisionEntity)
	if err != nil {
		_ = tx.Rollback()
		return divisionEntity, x.Wrap(err, "sql_create_division")
	}

	if err = tx.Commit(); err != nil {
		return divisionEntity, x.Wrap(err, "commit_create_division")
	}

	return divisionEntity, nil
}

func (d *division) GetDivisioByID(ctx context.Context, divisionID int64) (entity.Division, error) {
	result, err := d.getSQLDivisionByID(ctx, divisionID)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (d *division) GetDivision(ctx context.Context, cacheControl entity.CacheControl, param entity.DivisionParam) ([]entity.Division, entity.Pagination, error) {
	if cacheControl.MustRevalidate {
		result, pagination, err := d.getSQLDivision(ctx, param)
		if err != nil {
			return result, pagination, err
		}

		if err = d.setCacheDivision(ctx, param, result, pagination); err != nil {
			zerolog.Ctx(ctx).Warn().Err(err).Send()
		}

		return result, pagination, nil
	}

	result, pagination, err := d.getCacheDivision(ctx, param)
	if err == redis.Nil {
		zerolog.Ctx(ctx).Warn().Err(err).Send()

		result, pagination, err = d.getSQLDivision(ctx, param)
		if err != nil {
			return result, pagination, err
		}

		// save to cache
		if err = d.setCacheDivision(ctx, param, result, pagination); err != nil {
			zerolog.Ctx(ctx).Warn().Err(err).Send()
		}

		return result, pagination, nil
	} else if err != nil {
		zerolog.Ctx(ctx).Warn().Err(err).Send()

		// fallback if there is redis error e.g. bad conn, etc.
		// this is quite critical during high load traffic since it could be
		// thundering our db. (thundering herd).
		// we leave as it is to reduce code complexity [TODO LATER]
		return d.getSQLDivision(ctx, param)
	}

	return result, pagination, nil
}
