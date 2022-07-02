package middleware

import (
	"github.com/liansyyy/hade/framework/gin"
	"net/http"
)

// recovery机制，将协程中的函数异常进行捕获
func Recovery() gin.HandlerFunc {
	// 使用函数回调
	return func(c *gin.Context) {
		// 核心在增加这个recover机制，捕获c.Next()出现的panic
		defer func() {
			if err := recover(); err != nil {
				c.ISetStatus(http.StatusInternalServerError).IJson(err)
			}
		}()
		// 使用next执行具体的业务逻辑
		c.Next()
	}
}
