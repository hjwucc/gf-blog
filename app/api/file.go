package api

import (
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

func (a *apiFile) Put(r *ghttp.Request) {
	fileHeader := r.GetUploadFile("upload-file")
	file, err := fileHeader.Open()
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	// 如果存储目标为又拍云
	if target := r.GetRouterString("target"); "upy" == target {
		err = service.File.UpYunPut("/go-gf-blog/"+fileHeader.Filename, file.(io.Reader), map[string]string{"Content-Length": strconv.Itoa(int(fileHeader.Size))})
		if err != nil {
			response.JsonExit(r, 1, err.Error())
		}
	}
	response.JsonExit(r, 0, "上传文件成功")
}
