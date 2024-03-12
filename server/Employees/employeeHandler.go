package Employees

import (
	"errors"
	"example/company/database"
	"example/company/server"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GETEmployees(c *gin.Context) {
	db := database.RetrieveDatabase()

	rows, err := db.Query(database.GetAllEmployeesQuery)
	if err != nil {
		log.Println("Error querying employees: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: could not query database"})
		return
	}

	var employees []server.Employee
	for rows.Next() {
		var employee server.Employee
		var branch server.Branch
		if err := rows.Scan(&employee.EmpID, &employee.FirstName, &employee.LastName, &employee.BirthDate, &employee.Sex,
			&employee.Salary, &employee.SuperID, &branch.BranchID, &branch.BranchName,
			&branch.MgrID, &branch.MgrStartDate); err != nil {
			log.Println("Error retrieving employees", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: could not query database"})
			return
		}

		employee.Branch = branch

		employees = append(employees, employee)
	}
	c.IndentedJSON(http.StatusOK, employees)
}

func GETEmployee(c *gin.Context) {
	db := database.RetrieveDatabase()

	empID := c.Param("id")

	row := db.QueryRow(database.GetEmployeeByIDQuery, empID)

	var employee server.Employee
	if err := row.Scan(&employee.EmpID, &employee.FirstName, &employee.LastName, &employee.BirthDate, &employee.Sex, &employee.Salary, &employee.SuperID, &employee.Branch.BranchID); err != nil {
		log.Println("No Employee Found", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	c.IndentedJSON(http.StatusOK, employee)
}

//	func POSTEmployee(c *gin.Context) {
//		db := database.RetrieveDatabase()
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
func validateEmployeeData(employee server.Employee) error {
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
