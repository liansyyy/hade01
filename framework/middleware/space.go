package middleware

import (
	"github.com/liansyyy/hade/framework/gin"
	"log"
)

const Line = "-----------------------------------------------------------------\n\n"

func StartProcess() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Printf("\n\nrequest enter: %s\n", ctx.Request.RequestURI)
		ctx.Next()
		log.Printf("request exit: %s\n"+Line, ctx.Request.RequestURI)
	}
}
