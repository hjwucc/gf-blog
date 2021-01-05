package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"go-gf-blog/app/model"
	"go-gf-blog/app/service"
	"go-gf-blog/library/response"
	"strconv"
)

var Category = new(apiCategory)

type apiCategory struct {
}

// @summary 根据条件查找分类列表接口
// @tags    分类服务
// @produce json
// @param   entity  body model.ApiQueryCategoriesReq true "查找请求"
// @router  /category/conditionGetList [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiCategory) ConditionQueryList(r *ghttp.Request) {
	var apiReq *model.ApiQueryCategoriesReq
	var serviceReq *model.ServiceQueryCategoriesReq
	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	queryList, err := service.Category.ConditionQueryList(serviceReq)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "查询文章分类成功", g.Map{"categoryList": queryList})
}

// @summary 新增分类接口
// @tags    分类服务
// @produce json
// @param   entity  body model.ApiAddCategoryReq true "新增请求"
// @router  /category/add [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiCategory) Add(r *ghttp.Request) {
	var apiReq *model.ApiAddCategoryReq
	var serviceReq *model.ServiceAddCategoryReq
	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Category.Add(serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "文章分类保存成功", "success")
}

// @summary 刷新分类接口
// @tags    分类服务
// @produce json
// @router  /category/fresh [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiCategory) Fresh(r *ghttp.Request) {
	if err := service.Category.Fresh(); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "分类列表刷新成功", "success")
}

// @summary 修改分类接口
// @tags    分类服务
// @produce json
// @param   id path int true "分类ID"
// @router  /category/edit [PUT]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiCategory) Edit(r *ghttp.Request) {
	id, err := strconv.Atoi(r.GetRouterString("id"))
	if err != nil {
		response.JsonExit(r, 1, "文章分类id不正确")
	}
	var apiReq *model.ApiAddCategoryReq
	var serviceReq *model.ServiceAddCategoryReq
	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if err := service.Category.Edit(id, serviceReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "文章分类修改成功", "success")
}

// @summary 删除分类接口
// @tags    分类服务
// @produce json
// @param   id path int true "分类ID"
// @router  /category/delete [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiCategory) Delete(r *ghttp.Request) {
	id, err := strconv.Atoi(r.GetRouterString("id"))
	if err != nil {
		response.JsonExit(r, 1, "文章分类id不正确")
	}
	if err := service.Category.Delete(id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "文章分类删除成功")
}
