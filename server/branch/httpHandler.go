package branch

import (
	"example/company/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
