package auth

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes Auth routes
func ApplyRoutes(r *gin.RouterGroup) {
	auth := r.Group("/trickTips")
	{
		auth.POST("/", createUser)
		auth.POST("/", loginUser)
	}
}
