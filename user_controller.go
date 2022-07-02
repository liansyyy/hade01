package main

import (
	"fmt"
	"github.com/liansyyy/hade/framework/gin"
	"net/http"
	"time"
)

func UserLoginController(c *gin.Context) {
	foo, _ := c.DefaultQueryString("foo", "default value")
	time.Sleep(10 * time.Second)

	c.ISetStatus(http.StatusOK).
		IJson(fmt.Sprintf("ok, UserLoginController, uri: %s, foo: %s", c.Request.RequestURI, foo))

}
