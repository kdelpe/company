package employee

import (
	"errors"
	"example/company/database"
	"log"
)

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

	PUTEmployeeQuery = `
		UPDATE 
		    employee
		SET
			first_name = ?,
			last_name = ?,
			birth_date = ?,
			sex = ?,
			salary = ?,
			super_id = ?,
			branch_id = ?
		WHERE
			emp_id = ?
`
	DeleteEmployeeQuery = `
	DELETE FROM 
	           employee 
	       WHERE 
	           emp_id = ?
`
)

func updateEmployeeInDB(empID int64, employee Employee) error {
	db := database.RetrieveDatabase()

	// Prepare the SQL query to update the employee
	query := `
		UPDATE employee
		SET
			first_name = ?,
			last_name = ?,
			birth_date = ?,
			sex = ?,
			salary = ?,
			super_id = ?,
			branch_id = ?
		WHERE
			emp_id = ?
	`

	// Execute the update query
	result, err := db.Exec(query,
		employee.FirstName,
		employee.LastName,
		employee.BirthDate,
		employee.Sex,
		employee.Salary,
		employee.SuperID,
		employee.Branch.BranchID,
		empID,
	)
	if err != nil {
		log.Println("Error updating employee in database:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error retrieving number of rows affected:", err)
		return err
	}
	if rowsAffected == 0 {
		return errors.New("Update unsuccessful: Employee record not found")
	}

	return nil
}
