package api

import (
	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"go-gf-blog/app/auth"
	"go-gf-blog/app/model"
	"go-gf-blog/app/service"
	"go-gf-blog/library/response"
	"strconv"
)

// 文章API管理对象
var Article = new(apiArticle)

type apiArticle struct {
}

// @summary 根据文章ID查找接口
// @tags    文章服务
// @produce json
// @param   id path int true "文章ID"
// @router  /go-gf-blog/article/get [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiArticle) Get(r *ghttp.Request) {
	id, err := strconv.Atoi(r.GetRouterString("id"))
	if err != nil {
		response.JsonExit(r, 1, "文章id不正确")
	}
	entity, err := service.Article.Get(id)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "查询文章成功", entity)
}

// @summary 根据条件查找文章列表接口
// @tags    文章服务
// @produce json
// @param   entity  body model.ApiArticlesListReq true "查找请求"
// @router  /go-gf-blog/article/conditionGetList [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiArticle) ConditionGetList(r *ghttp.Request) {
	var apiReq *model.ApiArticlesListReq
	var serviceReq *model.ServiceArticlesListReq
	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	pageList, err := service.Article.ConditionPageList(serviceReq)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "根据条件查询文章成功", pageList)
}

// @summary 发布文章接口
// @tags    文章服务
// @produce json
// @param   entity  body model.ApiPublishArticleReq true "发布（新增、编辑）请求"
// @router  /go-gf-blog/article/publish [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiArticle) Publish(r *ghttp.Request) {
	var apiReq *model.ApiPublishArticleReq
	var serviceReq *model.ServicePublishArticleReq
	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Article.Publish(serviceReq, gconv.Int(auth.IdentityHandler(r).(jwt.MapClaims)["id"])); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "文章保存成功", "success")
}

// @summary 删除文章接口
// @tags    文章服务
// @produce json
// @param   id path int true "文章ID"
// @router  /go-gf-blog/article/delete [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiArticle) Delete(r *ghttp.Request) {
	id, err := strconv.Atoi(r.GetRouterString("id"))
	if err != nil {
		response.JsonExit(r, 1, "文章id不正确")
	}
	if err := service.Article.Delete(id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "文章删除成功", "success")

}

// @summary 修改文章属性(是否置顶，是否发布)接口
// @tags    文章服务
// @produce json
// @param   entity  body model.ApiUpdateArticleAttributeReq true "修改请求"
// @router  /go-gf-blog/article/updateAttributeById [PUT]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiArticle) UpdateAttributeById(r *ghttp.Request) {
	var apiReq *model.ApiUpdateArticleAttributeReq
	var serviceReq *model.ServiceUpdateArticleAttributeReq
	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Article.UpdateAttributeById(serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "文章属性修改成功", "success")
}
