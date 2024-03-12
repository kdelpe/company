package BranchSuppliers

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	router.GET("/api/branchsuppliers", GETBranchSuppliers)
	router.GET("/api/branchsuppliers/:id", GETBranchSupplier)
}
