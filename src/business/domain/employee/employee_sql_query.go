package employee

import (
	"strings"

	. "github.com/go-jet/jet/v2/mysql"
	. "github.com/linggaaskaedo/go-rocks/.gen/go-blog/table"

	"github.com/linggaaskaedo/go-rocks/src/business/entity"
)

func GetEmployeeByID(employeeID int64) (string, []interface{}) {
	stmt := SELECT(
		Employees.EmpNo,
		Employees.FirstName,
		Employees.LastName,
		Employees.Gender,
		Employees.HireDate,
	).FROM(
		Employees,
	).WHERE(
		Employees.EmpNo.EQ(Int(employeeID)),
	)

	return stmt.Sql()
}

func GetEmployee(param entity.EmployeeParam) (string, []interface{}, string, []interface{}) {
	var (
		stmt, stmtCount SelectStatement
		expID           []Expression
		exp             []BoolExpression
		orderClause     []OrderByClause
	)

	stmt = SELECT(
		Employees.EmpNo,
		Employees.FirstName,
		Employees.LastName,
		Employees.Gender,
	).FROM(
		Employees,
	)

	stmtCount = SELECT(
		COUNT(
			String("*"),
		),
	).FROM(
		Employees,
	)

	if param.ID != nil && len(param.ID) > 0 {
		for _, v := range param.ID {
			expID = append(expID, Int(v))
		}

		exp = append(exp, Employees.EmpNo.IN(expID...))
	}

	if param.FirstName != "" {
		if strings.Contains(param.FirstName, "%") {
			exp = append(exp, Employees.FirstName.LIKE(String(param.FirstName)))
		} else {
			exp = append(exp, Employees.FirstName.EQ(String(param.FirstName)))
		}

	}

	if param.LastName != "" {
		if strings.Contains(param.LastName, "%") {
			exp = append(exp, Employees.LastName.LIKE(String(param.LastName)))
		} else {
			exp = append(exp, Employees.LastName.EQ(String(param.LastName)))
		}
	}

	if param.Gender != "" {
		exp = append(exp, Employees.Gender.EQ(String(param.Gender)))
	}

	if exp != nil {
		stmt.WHERE(AND(exp...))
		stmtCount.WHERE(AND(exp...))
	}

	for _, order := range param.SortBy {
		switch ord := order; ord {
		// case "id":
		// 	orderClause = append(orderClause, Employees.EmpNo.ASC())
		// case "-id":
		// 	orderClause = append(orderClause, Employees.EmpNo.DESC())
		case "firstname":
			orderClause = append(orderClause, Employees.FirstName.ASC())
		case "-firstname":
			orderClause = append(orderClause, Employees.FirstName.DESC())
		case "lastname":
			orderClause = append(orderClause, Employees.LastName.ASC())
		case "-lastname":
			orderClause = append(orderClause, Employees.LastName.DESC())
		case "gender":
			orderClause = append(orderClause, Employees.Gender.CONCAT(StringColumn("gender")).ASC())
		case "-gender":
			orderClause = append(orderClause, Employees.Gender.CONCAT(StringColumn("gender")).DESC())
		default:
			orderClause = append(orderClause, Employees.EmpNo.ASC())
		}
	}

	stmt.ORDER_BY(orderClause...)
	stmt.OFFSET(param.Page)
	stmt.LIMIT(param.Limit)

	q1, a1 := stmt.Sql()
	q2, a2 := stmtCount.Sql()

	return q1, a1, q2, a2
}
