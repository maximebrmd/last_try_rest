package v1_0

import (
	"github.com/gin-gonic/gin"
	"last_try_rest/api/v1.0/auth"
	"last_try_rest/api/v1.0/trickTips"
	"last_try_rest/api/v1.0/user"
)

// ApplyRoutes godoc
func ApplyRoutes(rg *gin.RouterGroup) {
	v1 := rg.Group("/v1.0")
	{
		trickTips.ApplyRoutes(v1)
		user.ApplyRoutes(v1)
		auth.ApplyRoutes(v1)
	}
}
