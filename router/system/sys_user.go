package system

import (
	v1 "github.com/congwa/gin-start/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user")
	// 直接使用Router变量，而不是创建未使用的userRouter
	BaseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("resetPassword", BaseApi.ResetPassword)
	}
}
