package main

import (
	"example/company/database"
	"example/company/server/branch"
	"example/company/server/branch-suppliers"
	"example/company/server/client"
	"example/company/server/employee"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	database.Connection()

	router := gin.Default()

	// Set up routes
	branch.Routes(router)
	employee.Routes(router)
	branch_suppliers.Routes(router)
	client.Routes(router)

	//Start server
	if err := router.Run("localhost:9090"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
