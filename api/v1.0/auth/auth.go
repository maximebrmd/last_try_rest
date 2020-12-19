package auth

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes Auth routes
func ApplyRoutes(r *gin.RouterGroup) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", createUser)
		auth.POST("/login", loginUser)
	}
}
