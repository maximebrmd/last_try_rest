package trickTips

import (
	"github.com/gin-gonic/gin"
)

// ApplyRoutes trickTips routes
func ApplyRoutes(r *gin.RouterGroup) {
	trickTips := r.Group("/trickTips")
	{
		trickTips.POST("/", createTrickTips)
		trickTips.GET("/", getAllTrickTips)
		trickTips.PATCH("/:id/images", updateTrickTipsImages)
		//trickTips.PATCH("/:id", editTrickTips)
		//trickTips.DELETE("/:id", deleteTrickTips)
	}
}
