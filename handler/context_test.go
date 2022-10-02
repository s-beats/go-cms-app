package handler

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/s-beats/go-cms/middleware"
)

func TestSub(t *testing.T) {
	ctx := &gin.Context{}
	ctx.Set(string(middleware.KeySub), "test")
	t.Run("success",
		func(t *testing.T) {
			if got := Sub(ctx); got != "test" {
				t.Errorf("Sub() = %v, want Sub", got)
			}
		},
	)

	t.Run("success",
		func(t *testing.T) {
			casted := context.Context(ctx)
			val := casted.Value(string(middleware.KeySub))
			if val != "test" {
				t.Errorf("Sub() = %v, want Sub", val)
			}
		},
	)
}
