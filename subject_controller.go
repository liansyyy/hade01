package main

import (
	"codedemo01/framework"
	"fmt"
	"net/http"
)

func SubjectAddController(c *framework.Context) error {
	c.Json(http.StatusOK, fmt.Sprintf("ok, SubjectAddController, uri: %s", c.GetRequest().RequestURI))
	return nil
}

func SubjectListController(c *framework.Context) error {
	c.Json(http.StatusOK, fmt.Sprintf("ok, SubjectListController, uri: %s", c.GetRequest().RequestURI))
	return nil
}

func SubjectDelController(c *framework.Context) error {
	c.Json(http.StatusOK, fmt.Sprintf("ok, SubjectDelController, uri: %s", c.GetRequest().RequestURI))
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	c.Json(http.StatusOK, fmt.Sprintf("ok, SubjectUpdateController, uri: %s", c.GetRequest().RequestURI))
	return nil
}

func SubjectGetController(c *framework.Context) error {
	c.Json(http.StatusOK, fmt.Sprintf("ok, SubjectGetController, uri: %s", c.GetRequest().RequestURI))
	return nil
}

func SubjectNameController(c *framework.Context) error {
	c.Json(http.StatusOK, fmt.Sprintf("ok, SubjectNameController, uri: %s", c.GetRequest().RequestURI))
	return nil
}
