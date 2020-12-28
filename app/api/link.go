package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"go-gf-blog/app/model"
	"go-gf-blog/app/service"
	"go-gf-blog/library/response"
	"strconv"
)

// 链接API管理对象
var Link = new(apiLink)

type apiLink struct {
}

// @summary 根据条件查找链接列表接口
// @tags    链接服务
// @produce json
// @param   entity  body model.ApiLinkListReq true "查找请求"
// @router  /link/conditionGetList [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiLink) ConditionPageList(r *ghttp.Request) {
	var data *model.ApiLinkListReq
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	totalCount, pageList, err := service.Link.ConditionPageList(data)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	res := g.Map{
		"total":    totalCount,
		"pageList": pageList,
	}
	response.JsonExit(r, 0, "根据条件查询链接成功", res)
}

// @summary 新增链接接口
// @tags    链接服务
// @produce json
// @param   entity  body model.ApiAddLinkReq true "新增请求"
// @router  /link/add [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiLink) Add(r *ghttp.Request) {
	var data *model.ApiAddLinkReq
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if _, err := service.Link.Add(data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "链接添加成功", "success")
}

// @summary 修改链接接口
// @tags    链接服务
// @produce json
// @param   id path int true "链接ID"
// @router  /link/edit [PUT]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiLink) Edit(r *ghttp.Request) {
	id, err := strconv.Atoi(r.GetRouterString("id"))
	if err != nil {
		response.JsonExit(r, 1, "链接id不正确")
	}
	var data *model.ApiAddLinkReq
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if _, err := service.Link.Edit(id, data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "链接修改成功", "success")
}

// @summary 删除链接接口
// @tags    链接服务
// @produce json
// @param   id path int true "链接ID"
// @router  /link/delete [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiLink) Delete(r *ghttp.Request) {
	id, err := strconv.Atoi(r.GetRouterString("id"))
	if err != nil {
		response.JsonExit(r, 1, "链接id不正确")
	}
	if _, err := service.Link.Delete(id); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "链接删除成功", "success")
}
