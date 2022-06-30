package middleware

import (
	"codedemo01/framework"
	"log"
)

const Line = "-----------------------------------------------------------------\n\n"

func StartProcess() framework.ControllerHandler {
	return func(ctx *framework.Context) error {
		log.Printf("request enter: %s\n", ctx.GetRequest().RequestURI)
		ctx.Next()
		log.Printf("request exit: %s\n"+Line, ctx.GetRequest().RequestURI)
		return nil
	}
}
