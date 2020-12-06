package main

import (
	"github.com/gin-gonic/gin"
	"last_try_rest/api"
)

func main() {
	r := gin.Default()

	api.ApplyRoutes(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
