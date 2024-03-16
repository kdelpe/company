package client

import (
	"errors"
	"example/company/database"
	"example/company/server/branch"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GETClients(c *gin.Context) {
	db := database.RetrieveDatabase()

	rows, err := db.Query(GetAllClientsQuery)
	if err != nil {
		log.Println("Error retrieving client", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	var clients []Client
	for rows.Next() {
		var client Client
		var newBranch branch.Branch
		if err := rows.Scan(&newBranch.BranchID, &newBranch.BranchName, &newBranch.MgrID, &newBranch.MgrStartDate,
			&client.ClientID, &client.ClientName); err != nil {
			log.Println("Error retrieving client", err)
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		client.Branch = newBranch
		clients = append(clients, client)
	}
	c.IndentedJSON(http.StatusOK, clients)
}

func GETClient(c *gin.Context) {
	db := database.RetrieveDatabase()

	clientID, err := strconv.Atoi(c.Param("id")) // Convert string to integer
	if err != nil {
		log.Println("Invalid client ID:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}

	row := db.QueryRow(GetClientByIDQuery, clientID)

	var client Client
	var newBranch branch.Branch
	if err := row.Scan(&newBranch.BranchID, &newBranch.BranchName, &newBranch.MgrID, &newBranch.MgrStartDate,
		&client.ClientID, &client.ClientName); err != nil {
		log.Println("Error client", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
		return
	}
	client.Branch = newBranch
	c.IndentedJSON(http.StatusOK, client)
}

func POSTClient(c *gin.Context) {
	db := database.RetrieveDatabase()

	var newClient Client
	//validate the newClient
	if err := c.ShouldBindJSON(&newClient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := validateClient(newClient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "client_name field missing"})
		return
	}

	row, err := db.Exec(PostClientQuery, newClient.ClientName, newClient.Branch.BranchID)
	if err != nil {
		log.Println("Error adding new client:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Failed to add new client"})
		return
	}

	clientID, err := row.LastInsertId()
	if err != nil {
		log.Println("Failure to increment new employee ID", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newClient.ClientID = clientID

	c.IndentedJSON(http.StatusCreated, newClient)
}

func PUTClient(c *gin.Context) {
	db := database.RetrieveDatabase()

	clientID := parseParamID(c)

	var updatedClient Client
	if err := c.ShouldBindJSON(&updatedClient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec(PUTClientQuery, updatedClient.ClientName, updatedClient.Branch.BranchID, clientID)
	if err != nil {
		log.Println("Error updating client:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update client record"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Client record updated successfully"})
}

func DELETEClient(c *gin.Context) {
	db := database.RetrieveDatabase()

	clientID := parseParamID(c)

	var deletedClient Client
	if err := c.ShouldBindJSON(&deletedClient); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "client_name field is missing"})
		return
	}

	if deletedClient.ClientName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "client_name is missing"})
		return
	}
	//delete employee from the database
	_, err := db.Exec(DELETEClientQuery, clientID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete client"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Client record deleted successfully"})
}

func validateClient(client Client) error {
	if client.ClientName == "" {
		return errors.New("client name field is missing")
	}
	return nil
}

func parseParamID(c *gin.Context) int64 {
	clientID, err := strconv.ParseInt(c.Param("id"), 10, 64) // Convert string to integer
	if err != nil {
		log.Println("Invalid client ID:", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error:": "Invalid client ID"})
	}
	return clientID
}
