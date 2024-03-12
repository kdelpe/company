package Branches

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	router.GET("/api/branches", GETBranches)
	router.GET("/api/branches/:id", GETBranch)
}
