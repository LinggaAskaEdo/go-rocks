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

var Division = newDivisionTable("go-blog", "division", "")

type divisionTable struct {
	mysql.Table

	// Columns
	ID        mysql.ColumnInteger
	Name      mysql.ColumnString
	IsDeleted mysql.ColumnBool
	CreatedAt mysql.ColumnTimestamp
	UpdatedAt mysql.ColumnTimestamp
	DeletedAt mysql.ColumnTimestamp

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type DivisionTable struct {
	divisionTable

	NEW divisionTable
}

// AS creates new DivisionTable with assigned alias
func (a DivisionTable) AS(alias string) *DivisionTable {
	return newDivisionTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new DivisionTable with assigned schema name
func (a DivisionTable) FromSchema(schemaName string) *DivisionTable {
	return newDivisionTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new DivisionTable with assigned table prefix
func (a DivisionTable) WithPrefix(prefix string) *DivisionTable {
	return newDivisionTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new DivisionTable with assigned table suffix
func (a DivisionTable) WithSuffix(suffix string) *DivisionTable {
	return newDivisionTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newDivisionTable(schemaName, tableName, alias string) *DivisionTable {
	return &DivisionTable{
		divisionTable: newDivisionTableImpl(schemaName, tableName, alias),
		NEW:           newDivisionTableImpl("", "new", ""),
	}
}

func newDivisionTableImpl(schemaName, tableName, alias string) divisionTable {
	var (
		IDColumn        = mysql.IntegerColumn("id")
		NameColumn      = mysql.StringColumn("name")
		IsDeletedColumn = mysql.BoolColumn("is_deleted")
		CreatedAtColumn = mysql.TimestampColumn("created_at")
		UpdatedAtColumn = mysql.TimestampColumn("updated_at")
		DeletedAtColumn = mysql.TimestampColumn("deleted_at")
		allColumns      = mysql.ColumnList{IDColumn, NameColumn, IsDeletedColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn}
		mutableColumns  = mysql.ColumnList{NameColumn, IsDeletedColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn}
	)

	return divisionTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		Name:      NameColumn,
		IsDeleted: IsDeletedColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,
		DeletedAt: DeletedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}