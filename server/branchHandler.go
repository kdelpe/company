package server

import (
	"example/company/mysql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GETBranches(c *gin.Context) {
	db := mysql.RetrieveDatabase()

	rows, err := db.Query("SELECT * FROM branch;")
	if err != nil {
		log.Println("Error retrieving the branches: ", err)
		return
	}

	var branches []Branch
	for rows.Next() {
		var branch Branch
		if err := rows.Scan(&branch.BranchID, &branch.BranchName, &branch.MgrID, &branch.MgrStartDate); err != nil {
			log.Println("Error retrieving branches")
			return
		}
		branches = append(branches, branch)
	}
	c.IndentedJSON(http.StatusOK, branches)
}

func GETBranch(c *gin.Context) {
	db := mysql.RetrieveDatabase()

	branchID := c.Param("id")

	row := db.QueryRow("SELECT * FROM branch where branch_id = ?", branchID)

	var branch Branch
	if err := row.Scan(&branch.BranchID, &branch.BranchName, &branch.MgrID, &branch.MgrStartDate); err != nil {
		log.Println("Error retrieving branch")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}
	c.IndentedJSON(http.StatusOK, branch)
}
