package entity

import "github.com/linggaaskaedo/go-rocks/stdlib/preference/sqlx"

type Division struct {
	ID        int64         `db:"id"`
	Name      string        `db:"name"`
	IsDeleted bool          `db:"is_deleted"`
	CreatedAt sqlx.NullTime `db:"created_at"`
	UpdatedAt sqlx.NullTime `db:"updated_at"`
	DeletedAt sqlx.NullTime `db:"deleted_at"`
}

type DivisionParam struct {
	PublicID  []string      `param:"public_id" json:"public_id"`
	ID        []int64       `param:"id" db:"id"`
	Name      string        `param:"name" db:"name"`
	IsDeleted sqlx.NullBool `param:"is_deleted" db:"is_deleted"`
	SortBy    []string      `param:"sort_by" db:"sort_by"`
	Page      int64         `param:"page" db:"page"`
	Limit     int64         `param:"limit" db:"limit"`
}
