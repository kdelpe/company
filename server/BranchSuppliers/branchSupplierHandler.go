package BranchSuppliers

import (
	"example/company/database"
	"example/company/server"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GETBranchSuppliers(c *gin.Context) {
	db := database.RetrieveDatabase()

	rows, err := db.Query(database.GetAllBranchSuppliersQuery)
	if err != nil {
		log.Println("Error retrieving branch suppliers", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": "Bad Request"})
		return
	}

	var branchSuppliers []server.BranchSupplier
	for rows.Next() {
		var branchSupplier server.BranchSupplier
		var branch server.Branch
		if err := rows.Scan(&branch.BranchID, &branch.BranchName, &branch.MgrID, &branch.MgrStartDate,
			&branchSupplier.SupplierName, &branchSupplier.SupplyType); err != nil {
			log.Println("Error retrieving branch suppliers", err)
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": "Bad Request"})
			return
		}
		branchSupplier.Branch = branch
		branchSuppliers = append(branchSuppliers, branchSupplier)
	}
	c.IndentedJSON(http.StatusOK, branchSuppliers)
}

func GETBranchSupplier(c *gin.Context) {
	db := database.RetrieveDatabase()

	branchSupplierID, err := strconv.Atoi(c.Param("id")) // Convert string to integer
	if err != nil {
		log.Println("Invalid branch supplier ID:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": "Bad Request: Invalid branch supplier ID"})
		return
	}

	row := db.QueryRow(database.GetBranchSupplierByIDQuery, branchSupplierID)

	var branchSupplier server.BranchSupplier
	var branch server.Branch
	if err := row.Scan(&branch.BranchID, &branch.BranchName, &branch.MgrID, &branch.MgrStartDate,
		&branchSupplier.SupplierName, &branchSupplier.SupplyType); err != nil {
		log.Println("Error retrieving branch supplier", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": "Bad Request: Could not retrieve branch supplier"})
		return
	}
	branchSupplier.Branch = branch
	c.IndentedJSON(http.StatusOK, branchSupplier)
}
