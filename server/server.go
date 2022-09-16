package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/s-beats/go-cms/env"
	"github.com/s-beats/go-cms/factory"
	"github.com/s-beats/go-cms/infra"
	"github.com/s-beats/go-cms/middleware"
)

func Run() error {
	// ginのログ出さない
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	// Fixme: ややこしいのでMiddlewareのセットは一箇所にまとめたい
	db, err := infra.NewDB()
	if err != nil {
		return err
	}
	registroy := factory.NewRegistory(db, infra.NewS3Uploader())
	r.Use(middleware.RegistoryMiddleware(registroy))

	defineRoutes(r)

	hostname := env.HostName()
	port := env.Port()
	srv := &http.Server{
		Addr:    hostname + ":" + port,
		Handler: r,
	}

	fmt.Printf("listen %s \n", srv.Addr)
	go func() {
		// Start server
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println("listen and serve error")
			log.Fatal(err)
		}
	}()

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	sig := <-quit
	fmt.Printf("Shutdown Server with Signal %v\n", sig)
	// TODO: タイムアウトは環境変数から取得する
	timeout := time.Second * 10
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server Shutdown: %v\n", err)
	}
	fmt.Println("Server exiting")
	return nil
}
