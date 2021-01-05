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

	s.Group("/go-gf-blog", func(group *ghttp.RouterGroup) {
		group.ALL("/login",auth.GfJWTMiddleware.LoginHandler)
		group.ALL("/article/get/{id}",api.Article.Get)
		group.ALL("/article/conditionGetList",api.Article.ConditionGetList)
		group.ALL("/category/conditionGetList",api.Category.ConditionQueryList)
		group.ALL("/link/conditionGetList",api.Link.ConditionPageList)
		group.Middleware(service.Middleware.CORS,service.Middleware.MiddlewareAuth)
		group.ALL("/user/refreshToken",auth.GfJWTMiddleware.RefreshHandler)
		group.ALL("/user",api.User)
		group.ALL("/article/publish",api.Article.Publish)
		group.ALL("/article/delete/{id}",api.Article.Delete)
		group.ALL("/article/updateAttributeById",api.Article.UpdateAttributeById)
		group.ALL("/category/add",api.Category.Add)
		group.ALL("/category/delete/{id}",api.Category.Delete)
		group.ALL("/category/fresh/",api.Category.Fresh)
		group.ALL("/category/edit/{id}",api.Category.Edit)
		group.ALL("/link/add",api.Link.Add)
		group.ALL("/link/edit/{id}",api.Link.Edit)
		group.ALL("/link/delete/{id}",api.Link.Delete)
		group.ALL("/file/put/{target}",api.File.Put)
		group.ALL("/file/delete/{target}/{fileName}",api.File.Delete)
		group.ALL("/tag/add",api.Tag.Add)
		group.ALL("/tag/deleteTags",api.Tag.DeleteTags)
	})
}
