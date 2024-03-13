package client

import (
	"example/company/database"
	"example/company/server/branch"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GETClients(c *gin.Context) {
	db := database.RetrieveDatabase()

	rows, err := db.Query(database.GetAllClientsQuery)
	if err != nil {
		log.Println("Error retrieving client", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Bad Request: could not retrieve client"})
	}

	var clients []Client
	for rows.Next() {
		var client Client
		var branch branch.Branch
		if err := rows.Scan(&branch.BranchID, &branch.BranchName, &branch.MgrID, &branch.MgrStartDate,
			&client.ClientID, &client.ClientName); err != nil {
			log.Println("Error retrieving client", err)
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Bad Request: could not retrieve client"})
		}
		client.Branch = branch
		clients = append(clients, client)
	}
	c.IndentedJSON(http.StatusOK, clients)
}

func GETClient(c *gin.Context) {
	db := database.RetrieveDatabase()

	clientID, err := strconv.Atoi(c.Param("id")) // Convert string to integer
	if err != nil {
		log.Println("Invalid client ID:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": "Bad Request: Invalid client ID"})
		return
	}

	row := db.QueryRow(database.GetClientByIDQuery, clientID)

	var client Client
	var branch branch.Branch
	if err := row.Scan(&branch.BranchID, &branch.BranchName, &branch.MgrID, &branch.MgrStartDate,
		&client.ClientID, &client.ClientName); err != nil {
		log.Println("Error client", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": "Bad Request: Could not retrieve client"})
		return
	}
	client.Branch = branch
	c.IndentedJSON(http.StatusOK, client)
}