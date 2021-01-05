package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"go-gf-blog/app/model"
	"go-gf-blog/app/service"
	"go-gf-blog/library/response"
	"strconv"
)

var Tag = new(apiTag)

type apiTag struct {
}

// @summary 新增标签接口
// @tags    标签服务
// @produce json
// @param   entity  body model.ApiAddTagReq true "新增请求"
// @router  /go-gf-blog/tag/add [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiTag) Add(r *ghttp.Request) {
	var apiReq *model.ApiAddTagReq
	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if _, _, err := service.Tag.Add(apiReq.Name); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "标签添加成功", "success")
}

// @summary 删除标签接口
// @tags    标签服务
// @produce json
// @param   entity  body model.ApiDeleteTagsReq true "删除请求"
// @router  /go-gf-blog/tag/delete [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiTag) DeleteTags(r *ghttp.Request) {
	var apiReq *model.ApiDeleteTagsReq
	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	ids := gconv.SliceUint(apiReq.Ids)
	var successCount = len(ids)
	for _, v := range ids {
		err := service.Tag.Delete(int(v))
		if err != nil {
			successCount--
			glog.Error(err.Error())
		}
	}
	if successCount == 0 {
		response.JsonExit(r, 1, "没有需要删除的标签")
	}
	response.JsonExit(r, 0, "成功删除"+strconv.Itoa(successCount)+"个标签")
}
