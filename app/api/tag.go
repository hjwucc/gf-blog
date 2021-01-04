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

func (a *apiTag) Add(r *ghttp.Request) {
	var data *model.ApiAddTagReq
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	if _, _, err := service.Tag.Add(data.Name); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "标签添加成功", "success")
}

func (a *apiTag) DeleteTags(r *ghttp.Request) {
	var data *model.ApiDeleteTagsReq
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	ids := gconv.SliceUint(data.Ids)
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
