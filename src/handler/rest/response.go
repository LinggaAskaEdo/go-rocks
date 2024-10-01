package rest

import (
	"github.com/linggaaskaedo/go-rocks/src/business/dto"
	"github.com/linggaaskaedo/go-rocks/src/business/entity"
)

type HTTPErrResp struct {
	Meta dto.Meta `json:"metadata"`
}

type HTTPEmptyResp struct {
	Meta dto.Meta `json:"metadata"`
}

type HTTPDivisionResp struct {
	Meta dto.Meta     `json:"metadata"`
	Data DivisionData `json:"data"`
}

type DivisionData struct {
	Division *dto.DivisionDTO `json:"division,omitempty"`
}

type HTTPDivisionsResp struct {
	Meta       dto.Meta           `json:"metadata"`
	Data       DivisionsData      `json:"data"`
	Pagination *entity.Pagination `json:"pagination,omitempty"`
}

type DivisionsData struct {
	Divisions []dto.DivisionDTO `json:"divisions"`
}

type HTTPUserResp struct {
	Meta dto.Meta `json:"metadata"`
	Data UserData `json:"data"`
}

type UserData struct {
	User *dto.UserDTO `json:"user,omitempty"`
}

type HTTPUserLoginResp struct {
	Meta dto.Meta      `json:"metadata"`
	Data UserLoginData `json:"data"`
}

type UserLoginData struct {
	Token *dto.UserLoginDTO `json:"token,omitempty"`
}

type HTTPUserLogoutResp struct {
	Meta dto.Meta       `json:"metadata"`
	Data UserLogoutData `json:"data"`
}

type UserLogoutData struct {
	Token *dto.UserLogoutDTO `json:"data"`
}

type HTTPEmployeeResp struct {
	Meta dto.Meta     `json:"metadata"`
	Data EmployeeData `json:"data"`
}

type EmployeeData struct {
	Employee *dto.EmployeeDTO `json:"employee,omitempty"`
}

type HTTPKCEmployeeResp struct {
	Meta dto.Meta       `json:"metadata"`
	Data KCEmployeeData `json:"data"`
}

type KCEmployeeData struct {
	Employee *dto.KCEmployeeDTO `json:"employee,omitempty"`
}

type HTTPEmployeesResp struct {
	Meta       dto.Meta           `json:"metadata"`
	Data       EmployeesData      `json:"data"`
	Pagination *entity.Pagination `json:"pagination,omitempty"`
}

type EmployeesData struct {
	Employees []dto.EmployeeDTO `json:"employees"`
}
