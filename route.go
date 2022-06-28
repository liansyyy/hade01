package main

import "codedemo01/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
