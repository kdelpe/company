package branch_suppliers

import (
	"example/company/database"
	"example/company/server/branch"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GETBranchSuppliers(c *gin.Context) {
	db := database.RetrieveDatabase()

	rows, err := db.Query(GetAllBranchSuppliersQuery)
	if err != nil {
		log.Println("Error retrieving branch suppliers", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	var branchSuppliers []BranchSuppliers
	for rows.Next() {
		var branchSupplier BranchSuppliers
		var br branch.Branch
		if err := rows.Scan(&br.BranchID, &br.BranchName, &br.MgrID, &br.MgrStartDate,
			&branchSupplier.SupplierName, &branchSupplier.SupplyType); err != nil {
			log.Println("Error retrieving branch suppliers", err)
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
			return
		}
		branchSupplier.Branch = br
		branchSuppliers = append(branchSuppliers, branchSupplier)
	}
	c.IndentedJSON(http.StatusOK, branchSuppliers)
}

func GETBranchSupplier(c *gin.Context) {
	db := database.RetrieveDatabase()

	branchSupplierID, err := strconv.Atoi(c.Param("id")) // Convert string to integer
	if err != nil {
		log.Println("Invalid branch supplier ID:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	row := db.QueryRow(GetBranchSupplierByIDQuery, branchSupplierID)

	var branchSupplier BranchSuppliers
	var br branch.Branch
	if err := row.Scan(&br.BranchID, &br.BranchName, &br.MgrID, &br.MgrStartDate,
		&branchSupplier.SupplierName, &branchSupplier.SupplyType); err != nil {
		log.Println("Error retrieving branch supplier", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}
	branchSupplier.Branch = br
	c.IndentedJSON(http.StatusOK, branchSupplier)
}
