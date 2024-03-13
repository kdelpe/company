package server

type WorksWith struct {
	EmpID      *int64 `json:"emp_id"`
	ClientID   *int64 `json:"client_id"`
	TotalSales *int64 `json:"total_sales"`
}
