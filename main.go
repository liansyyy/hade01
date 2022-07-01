package main

import (
	"codedemo01/framework"
	"fmt"
	"net/http"
)

func main() {
	core := framework.NewCore()

	registerRouter(core)

	server := &http.Server{
		Handler: core,
		Addr:    ":8888",
	}

	server.ListenAndServe()

	fmt.Printf("111")
}
