package branch

import (
	"errors"
	"example/company/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GETBranches(c *gin.Context) {
	db := database.RetrieveDatabase()

	rows, err := db.Query(GetAllBranchesQuery)
	if err != nil {
		log.Println("Error retrieving the branch: ", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var branches []Branch
	for rows.Next() {
		var branch Branch
		if err := rows.Scan(&branch.BranchID, &branch.BranchName, &branch.MgrID, &branch.MgrStartDate); err != nil {
			log.Println("Error retrieving branch", err)
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		branches = append(branches, branch)
	}
	c.IndentedJSON(http.StatusOK, branches)
}

func GETBranch(c *gin.Context) {
	db := database.RetrieveDatabase()

	branchID := c.Param("id")

	row := db.QueryRow(GetBranchByIDQuery, branchID)

	var branch Branch
	if err := row.Scan(&branch.BranchID, &branch.BranchName, &branch.MgrID, &branch.MgrStartDate); err != nil {
		log.Println("Error retrieving branch", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, branch)
}

func POSTBranch(c *gin.Context) {
	db := database.RetrieveDatabase()

	var branch Branch
	if err := c.ShouldBindJSON(&branch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateBranch(branch); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	row, err := db.Exec(POSTBranchQuery, branch.BranchName, branch.MgrID, branch.MgrStartDate)
	if err != nil {
		log.Println("Error creating branch:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Failed to create branch"})
		return
	}

	branchID64, err := row.LastInsertId()
	if err != nil {
		log.Println("Failure to increment new employee ID", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	branchID := int8(branchID64)
	branch.BranchID = branchID

	c.IndentedJSON(http.StatusCreated, branch)
}

func PUTBranch(c *gin.Context) {
	db := database.RetrieveDatabase()

	branchID := parseParamID(c)

	var branch Branch
	if err := c.ShouldBindJSON(&branch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateBranch(branch); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec(PUTBranchQuery, branch.BranchName, branch.MgrID, branch.MgrStartDate, branchID)
	if err != nil {
		log.Println("Error updating branch:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update branch"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Branch updated successfully"})
}

func DELETEBranch(c *gin.Context) {
	db := database.RetrieveDatabase()

	branchID := parseParamID(c)

	var branch Branch
	if err := c.ShouldBindJSON(&branch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "branch_name is required to delete"})
		return
	}

	if err := validateBranch(branch); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if branch.BranchName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "branch_name is mandatory"})
		return
	}

	_, err := db.Exec(DELETEBranchQuery, branchID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete employee"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Employee deleted successfully"})

}

func validateBranch(branch Branch) error {
	if branch.BranchName == "" {
		return errors.New("branch name is required")
	}
	return nil
}

func parseParamID(c *gin.Context) int64 {
	empID, err := strconv.ParseInt(c.Param("id"), 10, 64) // Convert string to integer
	if err != nil {
		log.Println("Invalid branch ID:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
	}
	return empID
}
