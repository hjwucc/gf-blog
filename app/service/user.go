package service

import (
	"errors"
	"fmt"
	"go-gf-blog/app/dao"
	"go-gf-blog/app/model"
	"go-gf-blog/library/response"

	"github.com/gogf/gf/net/ghttp"
)

// 中间件管理服务
var User = new(serviceUser)

type serviceUser struct{}

// 用户注册
func (s *serviceUser) SignUp(r *model.ServiceUserSignUpReq) error {
	// 昵称为非必需参数，默认使用账号名称
	if r.Nickname == "" {
		r.Nickname = r.Passport
	}
	// 账号唯一性数据检查
	if !s.CheckPassport(r.Passport) {
		return errors.New(fmt.Sprintf("账号 %s 已经存在", r.Passport))
	}
	// 昵称唯一性数据检查
	if !s.CheckNickName(r.Nickname) {
		return errors.New(fmt.Sprintf("昵称 %s 已经存在", r.Nickname))
	}
	if _, err := dao.User.Save(r); err != nil {
		return err
	}
	return nil
}


// 用户注销
func (s *serviceUser) SignOut(r *ghttp.Request)  {
	// 用户登录时，将token返回给前端，前端本地存储，退出时，前端自己清理token
	response.JsonExit(r,0,"退出成功，记得清理token")
}

// 检查账号是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *serviceUser) CheckPassport(passport string) bool {
	if i, err := dao.User.FindCount("passport", passport); err != nil {
		return false
	} else {
		return i == 0
	}
}

// 检查昵称是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *serviceUser) CheckNickName(nickname string) bool {
	if i, err := dao.User.FindCount("nickname", nickname); err != nil {
		return false
	} else {
		return i == 0
	}
}

