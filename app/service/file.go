package service

import (
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"go-gf-blog/app/dao"
	"go-gf-blog/library/utils"
	"io"

	"github.com/upyun/go-sdk/v3/upyun"
)

var File = new(serviceFile)

type serviceFile struct {
}

var up *upyun.UpYun

func init() {
	// 查询配置参数
	config, err := dao.Config.FindOne("id", 1)
	if err != nil {
		err = gerror.New("初始化配置失败")
		return
	}
	// 初始化又拍云
	up = upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   config.YpuBucket,
		Operator: config.YpuOperator,
		Password: utils.AesDecrypt(config.YpyPassword, g.Cfg().Get("file.ypu_key").(string)),
	})
}

// 上传文件到又拍云
func (a *serviceFile) UpYunPut(path string, reader io.Reader, headers map[string]string) error {
	poc := upyun.PutObjectConfig{Path: path, Reader: reader, Headers: headers}
	if err := up.Put(&poc); err != nil {
		err = gerror.New("又拍云上传文件失败")
		return err
	}
	return nil
}
