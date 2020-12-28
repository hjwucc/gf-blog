package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
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
	var data *model.ApiQueryCategoriesReq
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	queryList, err := service.Category.ConditionQueryList(data)
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
	var data *model.ApiAddCategoryReq
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if _, err := service.Category.Add(data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "文章分类保存成功", "success")
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
	var data *model.ApiAddCategoryReq
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if _,err := service.Category.Edit(id,data); err != nil {
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
	if _, err := service.Category.Delete(id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "文章分类删除成功")
}
