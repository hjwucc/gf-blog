package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
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
	id := r.GetInt("id")
	entity, err := service.Article.Get(id)
	if err != nil {
		response.JsonExit(r, 1, "文章不存在,请联系管理员")
	}
	data := g.Map{
		"article": entity,
	}
	response.JsonExit(r, 0, "查询文章成功", data)
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
	total, pageList, err := service.Article.ConditionPageList(data)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	res := g.Map{
		"total":    total,
		"pageList": pageList,
	}
	response.JsonExit(r, 0, "根据条件查询文章成功", res)
}

// @summary 新增文章接口
// @tags    文章服务
// @produce json
// @param   entity  body model.ApiAddArticleReq true "新增请求"
// @router  /article/add [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiArticle) Add(r *ghttp.Request) {
	var data *model.ApiAddArticleReq
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if _, err := service.Article.Add(data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "文章保存成功", "success")
}

// @summary 修改文章接口
// @tags    文章服务
// @produce json
// @param   id path int true "文章ID"
// @router  /article/edit [PUT]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiArticle) Edit(r *ghttp.Request) {
	id, err := strconv.Atoi(r.GetRouterString("id"))
	if err != nil {
		response.JsonExit(r, 1, "文章id不正确")
	}
	var data *model.ApiAddArticleReq
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if _, err := service.Article.Edit(id, data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "文章修改成功", "success")
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
	if _, err := service.Article.Delete(id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "文章删除成功", "success")

}
