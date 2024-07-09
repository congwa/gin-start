package initialize

import (
	"github.com/congwa/gin-start/middleware"
	"github.com/congwa/gin-start/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	// 初始化路由
	Router := gin.New()
	Router.Use(gin.Recovery())
	Router.Use(gin.Logger())
	systemRouter := router.RouterGroupApp.System
	PublicGroup := Router.Group("v1")
	{
		systemRouter.InitBaseRouter(PublicGroup)
	}
	PrivateGroup := Router.Group("v1")
	PrivateGroup.Use(middleware.JWTAuth())
	{
		systemRouter.InitUserRouter(PrivateGroup)
	}
	return Router
}
