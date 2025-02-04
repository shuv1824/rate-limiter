package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shuv1824/rate-limiter/internal/domain"
)

func RateLimitMiddleware(limiter domain.RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		if strings.TrimSpace(clientIP) == "" {
			clientIP = "unknown"
		}

		if !limiter.Allow(clientIP) {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}

		c.Next()
	}
}
