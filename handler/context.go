package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/s-beats/go-cms/factory"
	"github.com/s-beats/go-cms/middleware"
)

func Registory(ctx *gin.Context) *factory.Registory {
	return ctx.MustGet(string(middleware.KeyRegistory)).(*factory.Registory)
}

func Sub(ctx *gin.Context) string {
	return ctx.MustGet(string(middleware.KeySub)).(string)
}
