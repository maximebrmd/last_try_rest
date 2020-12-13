package user

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes trickTips routes
func ApplyRoutes(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.POST("/:id/favorite", toggleFavorite)
		user.PATCH("/:id/avatar", updateAvatar)
		user.PATCH("/:id", updateUser)
	}
}
