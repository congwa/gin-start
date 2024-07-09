package core

import (
	time "time"

	"github.com/congwa/gin-start/global"
	"github.com/congwa/gin-start/initialize"
	"github.com/fvbock/endless"
)

// endless： Go 语言中平滑重启和无缝升级的 HTTP(S) 服务器包
// 平滑重启：服务器可以在不中断现有连接的情况下重启。
// 无缝升级：可以在不中断服务的情况下，升级服务器的二进制文件。
// 兼容性：与标准库的 net/http 完全兼容，易于集成到现有项目中。
// 简单易用：使用方式与标准的 HTTP 服务器类似，只需少量修改代码。

func RunServer() {
	initialize.Redis()
	Router := initialize.Routers()
	s := endless.NewServer(global.Config.Server.Host+":"+global.Config.Server.Port, Router)
	s.ReadHeaderTimeout = 20 * time.Second
	s.WriteTimeout = 20 * time.Second
	s.MaxHeaderBytes = 1 << 20

	global.LOG.Error(s.ListenAndServe().Error())
}
