package main

import (
	"codedemo01/framework"
	"net/http"
)

func UserLoginController(c *framework.Context) error {
	c.Json(http.StatusOK, "ok, UserLoginController")
	return nil
}
