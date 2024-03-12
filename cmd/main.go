package main

import (
	"example/company/database"
	"example/company/server/BranchSuppliers"
	"example/company/server/Branches"
	"example/company/server/Clients"
	"example/company/server/Employees"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	database.Connection()

	router := gin.Default()

	// Set up routes
	Branches.Routes(router)
	Employees.Routes(router)
	BranchSuppliers.Routes(router)
	Clients.Routes(router)

	//Start server
	if err := router.Run("localhost:9090"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
