//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/mysql"
)

var DeptEmp = newDeptEmpTable("go-blog", "dept_emp", "")

type deptEmpTable struct {
	mysql.Table

	// Columns
	EmpNo    mysql.ColumnInteger
	DeptNo   mysql.ColumnString
	FromDate mysql.ColumnDate
	ToDate   mysql.ColumnDate

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type DeptEmpTable struct {
	deptEmpTable

	NEW deptEmpTable
}

// AS creates new DeptEmpTable with assigned alias
func (a DeptEmpTable) AS(alias string) *DeptEmpTable {
	return newDeptEmpTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new DeptEmpTable with assigned schema name
func (a DeptEmpTable) FromSchema(schemaName string) *DeptEmpTable {
	return newDeptEmpTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new DeptEmpTable with assigned table prefix
func (a DeptEmpTable) WithPrefix(prefix string) *DeptEmpTable {
	return newDeptEmpTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new DeptEmpTable with assigned table suffix
func (a DeptEmpTable) WithSuffix(suffix string) *DeptEmpTable {
	return newDeptEmpTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newDeptEmpTable(schemaName, tableName, alias string) *DeptEmpTable {
	return &DeptEmpTable{
		deptEmpTable: newDeptEmpTableImpl(schemaName, tableName, alias),
		NEW:          newDeptEmpTableImpl("", "new", ""),
	}
}

func newDeptEmpTableImpl(schemaName, tableName, alias string) deptEmpTable {
	var (
		EmpNoColumn    = mysql.IntegerColumn("emp_no")
		DeptNoColumn   = mysql.StringColumn("dept_no")
		FromDateColumn = mysql.DateColumn("from_date")
		ToDateColumn   = mysql.DateColumn("to_date")
		allColumns     = mysql.ColumnList{EmpNoColumn, DeptNoColumn, FromDateColumn, ToDateColumn}
		mutableColumns = mysql.ColumnList{FromDateColumn, ToDateColumn}
	)

	return deptEmpTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		EmpNo:    EmpNoColumn,
		DeptNo:   DeptNoColumn,
		FromDate: FromDateColumn,
		ToDate:   ToDateColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
