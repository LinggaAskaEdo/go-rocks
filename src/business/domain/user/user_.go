package user

import (
	"context"
	"database/sql"

	"github.com/linggaaskaedo/go-rocks/src/business/entity"
	x "github.com/linggaaskaedo/go-rocks/stdlib/errors/entity"
)

func (u *user) CreateUser(ctx context.Context, userEntity entity.User) (entity.User, error) {
	tx, err := u.sql0.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
	})
	if err != nil {
		return userEntity, x.Wrap(err, "tx_create_user")
	}

	tx, userEntity, err = u.createSQLUser(ctx, tx, userEntity)
	if err != nil {
		_ = tx.Rollback()
		return userEntity, x.Wrap(err, "sql_create_user")
	}

	if err = tx.Commit(); err != nil {
		return userEntity, x.Wrap(err, "commit_create_user")
	}

	return userEntity, nil
}

func (u *user) GetUserByUserID(ctx context.Context, userID int64) (entity.User, error) {
	return u.getSQLUserByID(ctx, userID)
}

func (u *user) GetUserByUsername(ctx context.Context, username string) (entity.User, error) {
	return u.getSQLUserByUsername(ctx, username)
}
