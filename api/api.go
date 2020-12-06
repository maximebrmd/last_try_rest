package api

import (
	v1_0 "last_try_rest/api/v1.0"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes API routes
func ApplyRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		v1_0.ApplyRoutes(api)
	}
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
