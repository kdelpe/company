package branch

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	router.GET("/api/branch", GETBranches)
	router.GET("/api/branch/:id", GETBranch)
}
