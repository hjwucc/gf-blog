package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"go-gf-blog/app/auth"
	"go-gf-blog/app/service"
)

func init() {
	s := g.Server()
	// 分组路由注册方式
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/login",auth.GfJWTMiddleware.LoginHandler)
		group.Group("/user",func(group *ghttp.RouterGroup) {
			group.Middleware(service.Middleware.CORS,service.Middleware.MiddlewareAuth)
			group.ALL("/refresh_token", auth.GfJWTMiddleware.RefreshHandler)
		})
	})
}
