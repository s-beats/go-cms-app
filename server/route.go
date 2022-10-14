package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/s-beats/go-cms/handler"
	"github.com/s-beats/go-cms/middleware"
)

func defineRoutes(router *gin.Engine) {
	// health check
	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"health": "ok"})
	})

	setCommonMiddleware(router)

	v1 := router.Group("/v1")

	// 認証無し
	{
		v1.GET("/example", handler.Example)
	}
}

func setCommonMiddleware(router *gin.Engine) {
	router.Use(middleware.HandleError())
}
