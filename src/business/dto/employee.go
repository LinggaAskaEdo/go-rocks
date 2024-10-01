package dto

type EmployeeDTO struct {
	PublicID  string `json:"id" extensions:"x-order=0"`
	FirstName string `json:"firstname" extensions:"x-order=1"`
	LastName  string `json:"lastname" extensions:"x-order=2"`
	Gender    string `json:"gender" extensions:"x-order=3"`
}

type KCEmployeeDTO struct {
	ID        int64  `json:"id" extensions:"x-order=0"`
	FirstName string `json:"firstName" extensions:"x-order=1"`
	LastName  string `json:"lastName" extensions:"x-order=2"`
	Gender    string `json:"gender" extensions:"x-order=3"`
	HireDate  string `json:"hireDate" extensions:"x-order=4"`
}
