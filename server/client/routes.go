package client

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	router.GET("/api/clients/", GETClients)
	router.GET("/api/client/:id/", GETClient)
	router.POST("/api/client/", POSTClient)
	router.PUT("/api/client/:id", PUTClient)
	router.DELETE("/api/client/:id", DELETEClient)
}
