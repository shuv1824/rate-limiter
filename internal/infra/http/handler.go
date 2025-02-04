package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouter(rateLimiterMiddleware gin.HandlerFunc) *gin.Engine {
	r := gin.Default()

	r.Use(rateLimiterMiddleware)

	r.GET("/api/resources", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Request allowed"})
	})

	return r
}
