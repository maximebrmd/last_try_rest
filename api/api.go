package api

import (
	"github.com/gin-gonic/gin"
	v1_0 "last_try_rest/api/v1.0"
)

// ApplyRoutes API routes
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		v1_0.ApplyRoutes(api)
	}
}
