package server

import (
	"time"
)

type Employee struct {
	EmpID     *int64    `json:"emp_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date"`
	Sex       string    `json:"sex"`
	Salary    int64     `json:"salary"`
	SuperID   *int64    `json:"super_id"`
	Branch    Branch    `json:"branch"`
}

type Branch struct {
	BranchID     int8      `json:"branch_id"`
	BranchName   string    `json:"branch_name"`
	MgrID        int64     `json:"mgr_id"`
	MgrStartDate time.Time `json:"mgr_start_date"`
}

type BranchSupplier struct {
	Branch       Branch `json:"branch"`
	SupplierName string `json:"supplier_name"`
	SupplyType   string `json:"supply_type"`
}

type Client struct {
	ClientID   int64  `json:"client_id"`
	ClientName string `json:"client_name"`
	Branch     Branch `json:"branch"`
}

type WorksWith struct {
	EmpID      Employee
	ClientID   Client
	TotalSales int64 `json:"total_sales"`
}
