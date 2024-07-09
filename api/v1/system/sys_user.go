package system

import (
	"github.com/congwa/gin-start/global"
	"github.com/congwa/gin-start/model/common/response"
	"github.com/congwa/gin-start/model/system"
	systemReq "github.com/congwa/gin-start/model/system/request"
	systemRes "github.com/congwa/gin-start/model/system/response"
	"github.com/congwa/gin-start/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BaseApi struct{}

func (b *BaseApi) Login(c *gin.Context) {
	var l systemReq.Login

	err := c.ShouldBindJSON(&l)

	// key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	if l.CaptchaId != "" && l.Captcha != "" {
		u := &system.SysUser{Username: l.Username, Password: l.Password}
		user, err := userService.Login(u)
		if err != nil {
			response.FailWithMessage("用户名不存在或者密码错误", c)
			global.LOG.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
			return
		}
		// 签发jwt
		b.TokenNext(c, *user)
		return
	}
	response.FailWithMessage("验证码错误", c)
}

// TokenNext 登录以后签发jwt
func (b *BaseApi) TokenNext(c *gin.Context, user system.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.Config.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(systemReq.BaseClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		NickName: user.NickName,
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	response.OkWithDetailed(systemRes.LoginResponse{
		User:      user,
		Token:     token,
		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	}, "登录成功", c)
}

func (b *BaseApi) Register(c *gin.Context) {

	var r systemReq.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := &system.SysUser{Username: r.Username, NickName: r.NickName, Password: r.Password, Phone: r.Phone, Email: r.Email}
	userReturn, err := userService.Register(*user)
	if err != nil {
		global.LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册失败", c)
		return
	}
	response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
}

// ResetPassword 修改用户密码
func (b *BaseApi) ResetPassword(c *gin.Context) {
	var user system.SysUser
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = userService.ResetPassword(user.ID)
	if err != nil {
		global.LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}
