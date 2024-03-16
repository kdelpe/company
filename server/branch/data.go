package branch

type Branch struct {
	BranchID     int8   `json:"branch_id"`
	BranchName   string `json:"branch_name"`
	MgrID        int64  `json:"mgr_id"`
	MgrStartDate string `json:"mgr_start_date"`
}
