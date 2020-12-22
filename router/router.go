package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"go-gf-blog/app/api"
	"go-gf-blog/app/auth"
	"go-gf-blog/app/service"
)

func init() {
	s := g.Server()
	// 采用驼峰命名方式访问方法
	s.SetNameToUriType(ghttp.URI_TYPE_CAMEL)

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/login",auth.GfJWTMiddleware.LoginHandler)
		group.ALL("/article/get",api.Article.Get)
		group.ALL("/article/conditionGetList",api.Article.ConditionGetList)
		group.ALL("/category/conditionGetList",api.Category.ConditionQueryList)
		group.Middleware(service.Middleware.CORS,service.Middleware.MiddlewareAuth)
		group.ALL("/user/refreshToken",auth.GfJWTMiddleware.RefreshHandler)
		group.ALL("/user",api.User)
		group.ALL("/article/edit/{id}",api.Article.Edit)
		group.ALL("/article/delete/{id}",api.Article.Delete)
		group.ALL("/article/add",api.Article.Add)
		group.ALL("/category/add",api.Category.Add)
		group.ALL("/category/delete/{id}",api.Category.Delete)
		group.ALL("/category/edit/{id}",api.Category.Edit)
	})
}
