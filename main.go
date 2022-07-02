package main

import (
	"context"

	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/liansyyy/hade/framework/gin"
	"github.com/liansyyy/hade/framework/middleware"
)

func main() {
	core := gin.New()
	core.Use(middleware.StartProcess(), middleware.Cost(), middleware.Recovery())

	registerRouter(core)

	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}
	go func() {
		server.ListenAndServe()
	}()

	// 当前的goroutine等待信号量
	quit := make(chan os.Signal)
	// 监控信号：SIGINT, SIGTERM, SIGQUIT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	// 这里会阻塞当前goroutine等待信号
	<-quit

	withTimeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(withTimeoutCtx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

}
