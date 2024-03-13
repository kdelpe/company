package employee

import (
	"errors"
	"example/company/database"
	"example/company/server/branch"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GETEmployees(c *gin.Context) {
	db := database.RetrieveDatabase()

	rows, err := db.Query(database.GetAllEmployeesQuery)
	if err != nil {
		log.Println("Error querying employee: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request: could not query database"})
		return
	}

	var employees []Employee
	for rows.Next() {
		var employee Employee
		var branch branch.Branch
		if err := rows.Scan(&employee.EmpID, &employee.FirstName, &employee.LastName, &employee.BirthDate, &employee.Sex,
			&employee.Salary, &employee.SuperID, &branch.BranchID, &branch.BranchName,
			&branch.MgrID, &branch.MgrStartDate); err != nil {
			log.Println("Error retrieving employee", err)
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

	var employee Employee
	if err := row.Scan(&employee.EmpID, &employee.FirstName, &employee.LastName, &employee.BirthDate, &employee.Sex, &employee.Salary, &employee.SuperID, &employee.Branch.BranchID); err != nil {
		log.Println("No Employee Found", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	c.IndentedJSON(http.StatusOK, employee)
}

func POSTEmployee(c *gin.Context) {
	db := database.RetrieveDatabase()

	var employee Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateEmployeeData(employee); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	row, err := db.Exec(database.PostEmployeeQuery, employee.EmpID,
		employee.FirstName, employee.LastName, employee.BirthDate, employee.Sex, employee.Salary, employee.SuperID, employee.Branch.BranchID)
	if err != nil {
		log.Println("Error inserting a new employee", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	empID, err := row.LastInsertId()
	if err != nil {
		log.Println("Failure to increment new employee ID", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	employee.EmpID = &empID
	c.IndentedJSON(http.StatusCreated, employee)
}

func validateEmployeeData(employee Employee) error {
	if employee.FirstName == "" {
		return errors.New("first name is required")
	}

	if employee.LastName == "" {
		return errors.New("last name is required")
	}

	if employee.BirthDate == "" {
		return errors.New("birth date must be in the format YYYY-MM-DD")
	}

	if employee.Salary <= 50000 {
		return errors.New("salary must be greater than 50K")
	}

	if employee.Branch.BranchID <= 0 {
		return errors.New("employee's branch must be specified")
	}

	if employee.SuperID != employee.Branch.MgrID {
		return errors.New("incorrect supervisor id")
	}

	return nil
}
