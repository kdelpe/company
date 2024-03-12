package Clients

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	router.GET("/api/clients", GETClients)
	router.GET("/api/clients/:id", GETClient)
}
