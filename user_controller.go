package main

import (
	"codedemo01/framework"
	"fmt"
	"net/http"
)

func UserLoginController(c *framework.Context) error {
	c.Json(http.StatusOK, fmt.Sprintf("ok, UserLoginController, uri: %s", c.GetRequest().RequestURI))
	panic("panic by purpose")
	return nil
}
