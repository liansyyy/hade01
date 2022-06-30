package main

import (
	"codedemo01/framework"
	"codedemo01/framework/middleware"
)

// 注册路由规则
func registerRouter(core *framework.Core) {
	core.Use(middleware.StartProcess(), middleware.Cost(), middleware.Recovery())
	// 静态路由+HTTP方法匹配
	core.Get("/user/login", middleware.Test1(), UserLoginController)

	// 批量通用前缀
	subjectApi := core.Group("/subject")
	{
		subjectApi.Use(middleware.Test2())
		// 动态路由
		subjectApi.Delete("/:id", SubjectDelController)
		subjectApi.Put("/:id", SubjectUpdateController)
		subjectApi.Get("/:id", middleware.Test3(), SubjectGetController)
		subjectApi.Get("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Use(middleware.Test4())
			subjectInnerApi.Get("/name", SubjectNameController)
		}
	}
}
