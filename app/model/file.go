package model

import "io"

// 又拍云上传文件
type UpYunPutObjectConfig struct {
	Path              string             // 云存储中的路径
	LocalPath         string             // 待上传文件在本地文件系统中的路径
	Reader            io.Reader          // 待上传的内容
	Headers           map[string]string  // 额外的 HTTP 请求头
	UseMD5            bool               // 是否需要 MD5 校验
	UseResumeUpload   bool               // 是否使用断点续传
	AppendContent     bool               // 是否需要追加文件内容
	ResumePartSize    int64              // 断点续传块大小
	MaxResumePutTries int                // 断点续传最大重试次数
}

// 又拍云删除文件
type UpYunDeleteObjectConfig struct {
	Path  string        // 云存储中的路径
	Async bool          // 是否使用异步删除
}

// 又拍云初始化
type UpYunConfig struct {
	Bucket    string                // 云存储服务名（空间名）
	Operator  string                // 操作员
	Password  string                // 密码
	Secret    string                // 表单上传密钥，已经弃用！
	Hosts     map[string]string     // 自定义 Hosts 映射关系
	UserAgent string                // HTTP User-Agent 头，默认 "UPYUN Go SDK V2"
}
