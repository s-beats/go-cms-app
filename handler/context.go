package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/s-beats/go-cms/middleware"
)

func Sub(ctx *gin.Context) string {
	return ctx.MustGet(string(middleware.KeySub)).(string)
}
