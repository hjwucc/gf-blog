package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"go-gf-blog/app/model/articles"
	"go-gf-blog/app/service"
	"go-gf-blog/library/response"
	"strconv"
)

// 文章API管理对象
var Article = new(apiArticle)

type apiArticle struct {
}

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

func (a *apiArticle) ConditionGetList(r *ghttp.Request) {
	var data *articles.ApiArticlesListReq
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	total, pageList, err := service.Article.ConditionGetList(data)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	res := g.Map{
		"total":    total,
		"pageList": pageList,
	}
	response.JsonExit(r, 0, "根据条件查询文章成功", res)
}

func (a *apiArticle) Add(r *ghttp.Request) {
	var data *articles.ApiAddReq
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if _, err := service.Article.Add(data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "文章保存成功", "success")
}

func (a *apiArticle) Edit(r *ghttp.Request) {
	id, err := strconv.Atoi(r.GetRouterString("id"))
	if err != nil {
		response.JsonExit(r, 1, "文章id不正确")
	}
	var data *articles.ApiAddReq
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if _, err := service.Article.Edit(id, data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "文章修改成功", "success")
}

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
