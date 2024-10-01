package user

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"

	"github.com/linggaaskaedo/go-rocks/src/business/entity"
	commonerr "github.com/linggaaskaedo/go-rocks/stdlib/errors/common"
	x "github.com/linggaaskaedo/go-rocks/stdlib/errors/entity"
)

func (u *user) createSQLUser(ctx context.Context, tx *sqlx.Tx, userEntity entity.User) (*sqlx.Tx, entity.User, error) {
	row, err := tx.ExecContext(ctx, CreateUser, userEntity.Username, userEntity.Email, userEntity.Phone, userEntity.Division.ID, userEntity.Password, userEntity.CreatedAt)
	if err != nil {
		return tx, userEntity, x.Wrap(err, "create_user")
	}

	userEntity.ID, err = row.LastInsertId()
	zerolog.Ctx(ctx).Debug().Any("result", userEntity).Send()

	return tx, userEntity, nil
}

func (u *user) getSQLUserByID(ctx context.Context, userID int64) (entity.User, error) {
	result := entity.User{}

	row := u.sql0.QueryRowContext(ctx, GetUserByID, userID)

	if err := row.Scan(
		&result.ID,
		&result.Username,
		&result.Email,
		&result.Phone,
		&result.Password,
		&result.IsDeleted,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return result, x.WrapWithCode(err, commonerr.CodeSQLEmptyRow, "get_user_by_id")
		}

		return result, x.WrapWithCode(err, commonerr.CodeSQLRowScan, "get_user_by_id")
	}

	zerolog.Ctx(ctx).Debug().Any("result", result).Send()

	return result, nil
}

func (u *user) getSQLUserByUsername(ctx context.Context, username string) (entity.User, error) {
	result := entity.User{}

	row := u.sql0.QueryRowContext(ctx, GetUserByUsername, username)

	if err := row.Scan(
		&result.ID,
		&result.Username,
		&result.Email,
		&result.Phone,
		&result.Password,
		&result.IsDeleted,
		&result.CreatedAt,
		&result.UpdatedAt,
		&result.DeletedAt,
	); err != nil {
		if err == sql.ErrNoRows {
			return result, x.WrapWithCode(err, commonerr.CodeSQLEmptyRow, "get_user_by_username")
		}

		return result, x.WrapWithCode(err, commonerr.CodeSQLRowScan, "get_user_by_username")
	}

	zerolog.Ctx(ctx).Debug().Any("result", result).Send()

	return result, nil
}
