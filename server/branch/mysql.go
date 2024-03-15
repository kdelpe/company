package branch

const (
	GetAllBranchesQuery = `
SELECT 
	branch_id,
	branch_name,
	mgr_id,
	DATE_FORMAT(mgr_start_date, '%Y-%m-%d')
FROM 
	branch;
`

	GetBranchByIDQuery = `
SELECT 
		branch_id,
	    branch_name,
	    mgr_id,
	    DATE_FORMAT(mgr_start_date, '%Y-%m-%d')
FROM 
    branch 
WHERE 
    branch_id = ?;
`
)
