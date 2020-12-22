package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"go-gf-blog/app/model/categories"
	"go-gf-blog/app/service"
	"go-gf-blog/library/response"
	"strconv"
)

var Category = new(apiCategory)

type apiCategory struct {
}

func (a *apiCategory) ConditionQueryList(r *ghttp.Request) {
	var data *categories.ApiQueryReq
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	queryList, err := service.Category.ConditionQueryList(data)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "查询文章分类成功", g.Map{"categoryList": queryList})
}

func (a *apiCategory) Add(r *ghttp.Request) {
	var data *categories.ApiAddReq
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if _, err := service.Category.Add(data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "文章分类保存成功", "success")
}

func (a *apiCategory) Edit(r *ghttp.Request) {
	id, err := strconv.Atoi(r.GetRouterString("id"))
	if err != nil {
		response.JsonExit(r, 1, "文章分类id不正确")
	}
	var data *categories.ApiAddReq
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if _,err := service.Category.Edit(id,data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "文章分类修改成功", "success")
}

func (a *apiCategory) Delete(r *ghttp.Request) {
	id, err := strconv.Atoi(r.GetRouterString("id"))
	if err != nil {
		response.JsonExit(r, 1, "文章分类id不正确")
	}
	if _, err := service.Category.Delete(id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "文章分类删除成功")
}
