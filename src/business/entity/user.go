package entity

import "github.com/linggaaskaedo/go-rocks/stdlib/preference/sqlx"

type User struct {
	ID        int64         `db:"id"`
	Username  string        `db:"username"`
	Email     string        `db:"email"`
	Phone     string        `db:"phone"`
	Division  Division      `db:"division"`
	Password  string        `db:"password"`
	IsDeleted bool          `db:"is_deleted"`
	CreatedAt sqlx.NullTime `db:"created_at"`
	UpdatedAt sqlx.NullTime `db:"updated_at"`
	DeletedAt sqlx.NullTime `db:"deleted_at"`
}
