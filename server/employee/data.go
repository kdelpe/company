package employee

import (
	"example/company/server/branch"
)

type Employee struct {
	EmpID     *int64        `json:"emp_id"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	BirthDate string        `json:"birth_date"`
	Sex       string        `json:"sex"`
	Salary    int64         `json:"salary"`
	SuperID   *int64        `json:"super_id"`
	Branch    branch.Branch `json:"branch"`
}
