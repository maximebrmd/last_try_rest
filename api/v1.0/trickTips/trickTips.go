package trickTips

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes trickTips routes
func ApplyRoutes(r *gin.RouterGroup) {
	trickTips := r.Group("/trickTips")
	{
		trickTips.POST("/", addTrickTips)
		trickTips.GET("/", getAllTrickTips)
		//trickTips.PATCH("/:id", editTrickTips)
		//trickTips.DELETE("/:id", deleteTrickTips)
	}
}
