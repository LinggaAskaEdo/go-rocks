//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

// UseSchema sets a new schema name for all generated table SQL builder types. It is recommended to invoke
// this method only once at the beginning of the program.
func UseSchema(schema string) {
	Departments = Departments.FromSchema(schema)
	DeptEmp = DeptEmp.FromSchema(schema)
	DeptManager = DeptManager.FromSchema(schema)
	Division = Division.FromSchema(schema)
	Employees = Employees.FromSchema(schema)
	GorpMigrations = GorpMigrations.FromSchema(schema)
	Salaries = Salaries.FromSchema(schema)
	Titles = Titles.FromSchema(schema)
	User = User.FromSchema(schema)
}
