package main

import (
	"example/company/database"
	"example/company/server/branch"
	"example/company/server/branch-suppliers"
	"example/company/server/client"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	database.Connection()

	router := gin.Default()

	// Set up routes
	branch.Routes(router)
	Employees.Routes(router)
	branch_suppliers.Routes(router)
	client.Routes(router)

	//Start server
	if err := router.Run("localhost:9090"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
