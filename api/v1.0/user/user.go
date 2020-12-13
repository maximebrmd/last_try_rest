package user

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes trickTips routes
func ApplyRoutes(r *gin.RouterGroup) {
	user := r.Group("/user")
	{
		user.PATCH("/:id/avatar", updateAvatar)
		user.PATCH("/:id", updateUser)
		//trickTips.DELETE("/:id", deleteTrickTips)
	}
}
