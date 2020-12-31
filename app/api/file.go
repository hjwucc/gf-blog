package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"go-gf-blog/app/service"
	"go-gf-blog/library/response"
	"io"
	"strconv"
)

// 文件管理对象
var File = new(apiFile)

type apiFile struct {
}

// @summary 上传单个文件
// @tags    文件服务
// @produce json
// @param   upload-file body string true "文件"
// @param   target path string true "上传目标,如: upy/又拍云 qny/七牛云 aly/阿里云"
// @router  /file/put [PUT]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiFile) Put(r *ghttp.Request) {
	fileHeader := r.GetUploadFile("upload-file")
	file, err := fileHeader.Open()
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	// 如果存储目标为又拍云
	if target := r.GetRouterString("target"); "upy" == target {
		err, url := service.File.UpYunPut("/go-gf-blog/"+fileHeader.Filename, file.(io.Reader), map[string]string{"Content-Length": strconv.Itoa(int(fileHeader.Size))})
		if err != nil {
			response.JsonExit(r, 1, err.Error())
		}
		response.JsonExit(r, 0, "上传文件到又拍云成功", g.Map{"fileUrl": url})
	}
	response.JsonExit(r, 1, "暂未配置除又拍云外的其他存储")
}

// @summary 删除单个文件
// @tags    文件服务
// @produce json
// @param   target path string true "要删除的文件所在目标,如: upy/又拍云 qny/七牛云 aly/阿里云"
// @param   fileName path string true "要删除的文件名"
// @router  /file/delete [DELETE]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *apiFile) Delete(r *ghttp.Request) {
	// 要删除的文件存储在又拍云
	if target := r.GetRouterString("fileName"); "upy" == target {
		if err := service.File.UpYunDelete("/go-gf-blog/"+r.GetRouterString("fileName"), true); err != nil {
			response.JsonExit(r, 1, err.Error())
		}
		response.JsonExit(r, 0, "删除又拍云文件成功")
	}
	response.JsonExit(r, 1, "删除的文件不在又拍云，但暂未配置其他存储")
}
