package v1_0

import (
	"github.com/gin-gonic/gin"
	"last_try_rest/api/v1.0/trickTips"
)

// ApplyRoutes godoc
func ApplyRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/v1.0")
	{
		trickTips.ApplyRoutes(v1)
	}
}
