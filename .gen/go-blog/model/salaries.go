//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type Salaries struct {
	EmpNo    int32 `sql:"primary_key"`
	Salary   int32
	FromDate time.Time `sql:"primary_key"`
	ToDate   time.Time
}
