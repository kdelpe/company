package database

const (
	GetAllBranchesQuery = `SELECT * FROM branch;`

	GetBranchByIDQuery = "SELECT * FROM branch WHERE branch_id = ?;"
)
