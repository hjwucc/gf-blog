package service

import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"go-gf-blog/app/dao"
	"go-gf-blog/app/model"
	"go-gf-blog/library/utils"
	"io"

	"github.com/upyun/go-sdk/v3/upyun"
)

var File = new(serviceFile)

type serviceFile struct {
}

var up *upyun.UpYun
var config *model.Config

func init() {
	// 查询配置参数
	var err error
	config, err = dao.Config.FindOne("id", 1)
	if err != nil {
		err = gerror.New("初始化配置失败")
		return
	}
	// 初始化又拍云
	up = upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   config.UpyBucket,
		Operator: config.UpyOperator,
		Password: utils.AesDecrypt(config.UpyPassword, g.Cfg().Get("file.upy_key").(string)),
	})
}

// 上传文件到又拍云,返回文件url
func (a *serviceFile) UpYunPut(path string, reader io.Reader, headers map[string]string) (err error, url string) {
	poc := upyun.PutObjectConfig{Path: path, Reader: reader, Headers: headers}
	if err := up.Put(&poc); err != nil {
		err = gerror.New("又拍云上传文件失败")
		return err, ""
	}
	return nil, config.UpyUrl + path
}

// 又拍云删除文件
func (a *serviceFile) UpYunDelete(path string, async bool) error {
	if err := up.Delete(&upyun.DeleteObjectConfig{
		Path:   path,
		Async:  async,
		Folder: false,
	}); err != nil {
		err = gerror.New("又拍云删除文件失败")
		return err
	}
	return nil
}
