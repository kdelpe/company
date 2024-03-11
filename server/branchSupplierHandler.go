package server

import (
	"example/company/mysql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func GETBranchSuppliers(c *gin.Context) {
	db := mysql.RetrieveDatabase()

	query := `
	SELECT
		b.branch_id,
		b.branch_name,
		b.mgr_id,
		b.mgr_start_date,
		bs.supplier_name,
		bs.supply_type
	FROM
		branch b
	LEFT JOIN
    	branch_supplier bs ON b.branch_id = bs.branch_id;
`
	rows, err := db.Query(query)
	if err != nil {
		log.Println("Error retrieving branch suppliers", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": "Bad Request"})
		return
	}

	var branchSuppliers []BranchSupplier
	for rows.Next() {
		var branchSupplier BranchSupplier
		var branch Branch
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
