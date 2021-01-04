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
// @router  /article/get [GET]
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
// @router  /article/conditionGetList [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiArticle) ConditionGetList(r *ghttp.Request) {
	var data *model.ApiArticlesListReq
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	pageList, err := service.Article.ConditionPageList(data)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	response.JsonExit(r, 0, "根据条件查询文章成功", pageList)
}

// @summary 发布文章接口
// @tags    文章服务
// @produce json
// @param   entity  body model.ApiAddArticleReq true "发布（新增、编辑）请求"
// @router  /article/add [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiArticle) Publish(r *ghttp.Request) {
	var data *model.ApiPublishArticleReq
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	if err := service.Article.Publish(data, gconv.Int(auth.IdentityHandler(r).(jwt.MapClaims)["id"])); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "文章保存成功", "success")
}

// @summary 删除文章接口
// @tags    文章服务
// @produce json
// @param   id path int true "文章ID"
// @router  /article/delete [DELETE]
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
