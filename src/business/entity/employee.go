package entity

import "github.com/linggaaskaedo/go-rocks/stdlib/preference/sqlx"

type Employee struct {
	EmpNo     int64         `db:"emp_no"`
	FirstName string        `db:"first_name"`
	LastName  string        `db:"last_name"`
	Gender    string        `db:"gender"`
	HireDate  sqlx.NullTime `db:"hire_date"`
}

type EmployeeParam struct {
	PublicID  []string `param:"public_id" json:"public_id"`
	ID        []int64  `param:"id" db:"id"`
	FirstName string   `param:"first_name" db:"first_name"`
	LastName  string   `param:"last_name" db:"last_name"`
	Gender    string   `param:"gender" db:"gender"`
	SortBy    []string `param:"sort_by" db:"sort_by"`
	Page      int64    `param:"page" db:"page"`
	Limit     int64    `param:"limit" db:"limit"`
}
