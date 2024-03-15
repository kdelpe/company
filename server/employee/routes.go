package employee

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	router.GET("/api/employee", GETEmployees)
	router.GET("/api/employee/:id", GETEmployee)
	router.POST("/api/employee", POSTEmployee)
	router.PUT("/api/employee/:id", PUTEmployee)
}
