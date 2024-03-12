package Employees

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	router.GET("/api/employees", GETEmployees)
	router.GET("/api/employees/:id", GETEmployee)
}
