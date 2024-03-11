package main

import (
	"example/company/mysql"
	"example/company/server"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	mysql.Connection()

	router := gin.Default()

	router.GET("/branches", server.GETBranches)
	router.GET("/branches/:id", server.GETBranch)

	router.GET("/employees", server.GETEmployees)
	router.GET("/employees/:id", server.GETEmployee)

	router.GET("/branchsuppliers", server.GETBranchSuppliers)

	//Start server
	err := router.Run("localhost:9090")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
