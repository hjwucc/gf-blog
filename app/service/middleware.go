package service

import (
	"github.com/gogf/gf/net/ghttp"
	"go-gf-blog/app/auth"
)

// 中间件管理服务
var Middleware = new(serviceMiddleware)

type serviceMiddleware struct{}


// 允许接口跨域请求
func (s *serviceMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *serviceMiddleware) MiddlewareAuth(r *ghttp.Request) {
	auth.GfJWTMiddleware.MiddlewareFunc()(r)
	r.Middleware.Next()
}

