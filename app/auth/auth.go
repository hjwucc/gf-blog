package auth

import (
	"errors"
	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/crypto/gsha1"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gvalid"
	"go-gf-blog/app/dao"
	"go-gf-blog/library/utils"
	"time"
)

var (
	// The underlying JWT middleware.
	GfJWTMiddleware *jwt.GfJWTMiddleware
	// Customized login parameter validation rules.
	ValidationRules = g.Map{
		"username": "required",
		"password": "required",
	}
)

// Initialization function,
// rewrite this function to customized your own JWT settings.
func init() {
	authMiddleware, err := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "test zone",
		Key:             []byte("secret key"),
		Timeout:         time.Minute * 30,
		MaxRefresh:      time.Minute * 30,
		IdentityKey:     "id",
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:   "Bearer",
		TimeFunc:        time.Now,
		Authenticator:   Authenticator,
		LoginResponse:   LoginResponse,
		RefreshResponse: RefreshResponse,
		Unauthorized:    Unauthorized,
		IdentityHandler: IdentityHandler,
		PayloadFunc:     PayloadFunc,
	})
	if err != nil {
		glog.Fatal("JWT Error:" + err.Error())
	}
	GfJWTMiddleware = authMiddleware
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler sets the identity for JWT.
func IdentityHandler(r *ghttp.Request) interface{} {
	claims := jwt.ExtractClaims(r)
	return claims["nickname"]
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(r *ghttp.Request, code int, message string) {
	r.Response.WriteJson(g.Map{
		"code": 0,
		"msg":  message,
	})
	r.ExitAll()
}

// LoginResponse is used to define customized login-successful callback function.
func LoginResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "登录成功",
		"data": g.Map{
			"token":    token,
			"expire":   expire.Format(time.RFC3339),
			"nickname": r.GetParam("nickname"),
		},
	})
	r.ExitAll()
}

// RefreshResponse is used to get a new token no matter current token is expired or not.
func RefreshResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "刷新token成功",
		"data": g.Map{
			"token":    token,
			"expire":   expire.Format(time.RFC3339),
			"nickname": IdentityHandler(r),
		},
	})
	r.ExitAll()
}

// Authenticator is used to validate login parameters.
// It must return user data as user identifier, it will be stored in Claim Array.
// Check error (e) to determine the appropriate error message.
func Authenticator(r *ghttp.Request) (interface{}, error) {
	data := r.GetMap()
	if e := gvalid.CheckMap(data, ValidationRules); e != nil {
		return "", jwt.ErrFailedAuthentication
	}
	// 从数据库读取用户数据并验证
	user, err := dao.User.FindOne("passport=? and password=?", data["passport"], gsha1.Encrypt(data["password"]))
	if user == nil || err != nil {
		return "", errors.New("用户名或密码错误")
	}
	// 更新用户最后登录时间和IP
	res, _ := dao.User.Data(g.Map{"last_login_time": gtime.Now(), "last_login_ip": utils.RemoteIp(r)}).Where("id", user.Id).Update()
	if affected, errs := res.RowsAffected(); affected == 0 || errs != nil {
		glog.Error("刷新用户最后登录状态失败")
	}
	r.SetParam("nickname", user.Nickname)
	return g.Map{
		"nickname": user.Nickname,
		"id":       data["passport"],
	}, nil
}
