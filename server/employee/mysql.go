package employee

const (
	GetAllEmployeesQuery = `
			SELECT 
				employee.emp_id, 
				employee.first_name, 
				employee.last_name, 
				DATE_FORMAT(employee.birth_date, '%Y-%m-%d'), 
				employee.sex, 
				employee.salary, 
				employee.super_id, 
				branch.branch_id, 
				branch.branch_name, 
				branch.mgr_id, 
				DATE_FORMAT(branch.mgr_start_date, '%Y-%m-%d') 
			FROM 
				employee
			LEFT JOIN 
				branch ON employee.branch_id = branch.branch_id
			LEFT JOIN 
				employee mgr ON branch.mgr_id = mgr.emp_id
			LEFT JOIN 
				employee super ON employee.super_id = super.emp_id;
		`

	GetEmployeeByIDQuery = `
			SELECT 
				employee.emp_id, 
				employee.first_name, 
				employee.last_name, 
				DATE_FORMAT(employee.birth_date, '%Y-%m-%d'), 
				employee.sex, 
				employee.salary, 
				employee.super_id, 
				branch.branch_id, 
				branch.branch_name, 
				branch.mgr_id, 
				DATE_FORMAT(branch.mgr_start_date, '%Y-%m-%d') 
			FROM 
				employee
			LEFT JOIN 
				branch ON employee.branch_id = branch.branch_id
			LEFT JOIN 
				employee mgr ON branch.mgr_id = mgr.emp_id
			LEFT JOIN 
				employee super ON employee.super_id = super.emp_id 
			WHERE 
			    employee.emp_id = ?;`

	PostEmployeeQuery = `INSERT INTO employee VALUES (?, ?, ?, ?, ?, ?, ?, ?);`
)
