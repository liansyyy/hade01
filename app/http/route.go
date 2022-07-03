package http

import (
	"github.com/liansyyy/hade/app/http/module/demo"
	"github.com/liansyyy/hade/framework/gin"
)

func Routes(r *gin.Engine) {

	r.Static("/dist/", "./dist/")

	demo.Register(r)
}
