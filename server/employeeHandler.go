package server

import (
	"errors"
	"example/company/mysql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GETEmployees(c *gin.Context) {
	db := mysql.RetrieveDatabase()

	query := `
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
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error querying employees: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	var employees []Employee
	for rows.Next() {
		var employee Employee
		var branch Branch
		if err := rows.Scan(&employee.EmpID, &employee.FirstName, &employee.LastName, &employee.BirthDate, &employee.Sex,
			&employee.Salary, &employee.SuperID, &branch.BranchID, &branch.BranchName,
			&branch.MgrID, &branch.MgrStartDate); err != nil {
			log.Println("Error retrieving employees", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		employee.Branch = branch

		employees = append(employees, employee)
	}
	c.IndentedJSON(http.StatusOK, employees)
}

func GETEmployee(c *gin.Context) {
	db := mysql.RetrieveDatabase()

	empID := c.Param("id")

	row := db.QueryRow("SELECT * FROM employee WHERE emp_id = ?", empID)

	var employee Employee
	if err := row.Scan(&employee.EmpID, &employee.FirstName, &employee.LastName, &employee.BirthDate, &employee.Sex, &employee.Salary, &employee.SuperID, &employee.Branch.BranchID); err != nil {
		log.Println("No Employee Found", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	c.IndentedJSON(http.StatusOK, employee)
}

//	func POSTEmployee(c *gin.Context) {
//		db := mysql.RetrieveDatabase()
//
//		var employee Employee
//		if err := c.ShouldBindJSON(&employee); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//
//		row, err := db.Exec("INSERT INTO employee VALUES(?, ?, ?, ?, ?, ?, ?, ?)", employee.EmpID,
//			employee.FirstName, employee.LastName, employee.BirthDate, employee.Sex, employee.Salary, employee.SuperID, employee.Branch.BranchID)
//		if err != nil {
//			log.Println("Error inserting a new employee", err)
//			return
//		}
//
//		empID, err := row.LastInsertId()
//		if err != nil {
//			log.Println("Error retrieving new employee ID", err)
//			return
//		}
//
//		employee.EmpID = empID
//		c.IndentedJSON(http.StatusCreated, employee)
//	}
func validateEmployeeData(employee Employee) error {
	if employee.FirstName == "" {
		return errors.New("first name is required")
	}

	if employee.LastName == "" {
		return errors.New("last name is required")
	}

	if employee.BirthDate.IsZero() {
		return errors.New("birth date is required")
	}

	if employee.Salary <= 0 {
		return errors.New("salary must be greater than zero")
	}

	if employee.Branch.BranchID <= 0 {
		return errors.New("employee's branch must be specified")
	}

	return nil
}
