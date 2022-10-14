package middleware

import (
	"github.com/gin-gonic/gin"
)

// TODO
func SetupCorsConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
