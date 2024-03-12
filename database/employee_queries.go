package database

const (
	GetAllEmployeesQuery = `
	    SELECT 
            employee.emp_id, 
            employee.first_name, 
            employee.last_name, 
            employee.birth_date, 
            employee.sex, 
            employee.salary, 
            employee.super_id, 
            branch.branch_id, 
            branch.branch_name, 
            branch.mgr_id, 
            branch.mgr_start_date 
        FROM 
            employee employee
        LEFT JOIN 
            branch branch ON employee.branch_id = branch.branch_id
        LEFT JOIN 
            employee mgr ON branch.mgr_id = mgr.emp_id
        LEFT JOIN 
            employee super ON employee.super_id = super.emp_id;
	`

	GetEmployeeByIDQuery = "SELECT * FROM employee WHERE emp_id = ?"
)
