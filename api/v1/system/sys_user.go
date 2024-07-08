package system

import (
	"github.com/congwa/gin-start/global"
	"github.com/congwa/gin-start/model/common/response"
	"github.com/congwa/gin-start/model/system"
	systemReq "github.com/congwa/gin-start/model/system/request"
	systemRes "github.com/congwa/gin-start/model/system/response"

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
		// TODO: 进行登录的下一步 中间件处理
		return
	}
	response.FailWithMessage("验证码错误", c)
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
