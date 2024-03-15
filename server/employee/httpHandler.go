package employee

import (
	"errors"
	"example/company/database"
	"example/company/server/branch"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GETEmployees(c *gin.Context) {
	db := database.RetrieveDatabase()

	rows, err := db.Query(GetAllEmployeesQuery)
	if err != nil {
		log.Println("Error querying employee: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var employees []Employee
	for rows.Next() {
		var employee Employee
		var br branch.Branch
		if err := rows.Scan(&employee.EmpID, &employee.FirstName, &employee.LastName, &employee.BirthDate, &employee.Sex,
			&employee.Salary, &employee.SuperID, &br.BranchID, &br.BranchName,
			&br.MgrID, &br.MgrStartDate); err != nil {
			log.Println("Error retrieving employee", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		employee.Branch = br
		employee.SuperID = br.MgrID
		employees = append(employees, employee)
	}
	c.IndentedJSON(http.StatusOK, employees)
}

func GETEmployee(c *gin.Context) {
	db := database.RetrieveDatabase()

	empID := parseParamID(c)

	row := db.QueryRow(GetEmployeeByIDQuery, empID)

	var employee Employee
	var br branch.Branch
	if err := row.Scan(&employee.EmpID, &employee.FirstName, &employee.LastName, &employee.BirthDate,
		&employee.Sex, &employee.Salary, &employee.SuperID, &br.BranchID, &br.BranchName,
		&br.MgrID, &br.MgrStartDate); err != nil {
		log.Println("No Employee Found", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	employee.Branch = br
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

	row, err := db.Exec(PostEmployeeQuery, employee.EmpID, employee.FirstName, employee.LastName,
		employee.BirthDate, employee.Sex, employee.Salary, &employee.SuperID, employee.Branch.BranchID)
	if err != nil {
		log.Println("Error inserting a new employee", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if employee.SuperID == nil {
		employee.SuperID = employee.Branch.MgrID
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

func PUTEmployee(c *gin.Context) {

	empID := int64(parseParamID(c))

	var updatedEmployee Employee
	if err := c.ShouldBindJSON(&updatedEmployee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateEmployeeData(updatedEmployee); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := updateEmployeeInDB(empID, updatedEmployee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Employee updated successfully"})
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

	return nil
}

func parseParamID(c *gin.Context) int {
	empID, err := strconv.Atoi(c.Param("id")) // Convert string to integer
	if err != nil {
		log.Println("Invalid employee ID:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
	}
	return empID
}
