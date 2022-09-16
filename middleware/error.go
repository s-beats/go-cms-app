package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// https://github.com/gin-gonic/gin/issues/274
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		errorToPrint := c.Errors.ByType(gin.ErrorTypePrivate).Last()
		if errorToPrint != nil {
			log.Err(errorToPrint).Send()
			c.AbortWithStatusJSON(
				http.StatusBadRequest,
				gin.H{"error": errorToPrint},
			)
		}
	}
}
