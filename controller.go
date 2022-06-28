package main

import (
	"codedemo01/framework"
	"context"
	"log"
	"net/http"
	"time"
)

func FooControllerHandler(ctx *framework.Context) error {
	log.Printf("dealing: %s\n", ctx.GetRequest().RequestURI)
	finish := make(chan struct{}, 1)
	paincChan := make(chan interface{}, 1)

	withTimeoutCtx, cancel := context.WithTimeout(ctx.BaseContext(), 5*time.Second)
	defer cancel()

	go func() {
		defer func() {
			if p := recover(); p != nil {
				paincChan <- p
			}
		}()
		time.Sleep(5 * time.Second)
		finish <- struct{}{}
	}()

	select {
	//監聽Panic
	case p := <-paincChan:
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		log.Println(p)
		ctx.Json(http.StatusInternalServerError, "panic")
	//監聽結束事件
	case <-finish:
		ctx.Json(http.StatusOK, "ok")
		log.Printf("request: %s, finish!\n", ctx.GetRequest().RequestURI)
		log.Printf("request: %s, finish!\n", ctx.GetRequest().URL.Path)
	//監聽超時事件
	case <-withTimeoutCtx.Done():
		ctx.WriterMux().Lock()
		defer ctx.WriterMux().Unlock()
		ctx.Json(http.StatusInternalServerError, "time out")
		ctx.SetHasTimeout()
	}
	return nil
}
