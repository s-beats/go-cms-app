package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/s-beats/go-cms/factory"
)

func RegistoryMiddleware(registory *factory.Registory) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(string(KeyRegistory), registory)
		c.Next()
	}
}
