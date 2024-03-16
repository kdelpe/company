package branch

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine) {
	router.GET("/api/branch", GETBranches)
	router.GET("/api/branch/:id", GETBranch)
	router.POST("api/branch", POSTBranch)
	router.PUT("api/branch/:id", PUTBranch)
	router.DELETE("api/branch/:id", DELETEBranch)
}
