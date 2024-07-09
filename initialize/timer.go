package initialize

import (
	"fmt"

	"github.com/congwa/gin-start/global"
	"github.com/congwa/gin-start/model/system"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

func Timer() {
	go func() {
		var option []cron.Option
		option = append(option, cron.WithSeconds())

		// TODO： 进行一个定时查库的任务
		_, err := global.Timer.AddTaskByFunc("GetUserCount", "@daily", func() {
			var UserCount int64
			if err := global.DB.Model(&system.SysUser{}).Count(&UserCount).Error; err != nil {
				global.LOG.Error(err.Error())
			}

			// 因为Count方法直接修改UserCount，所以我们只需要记录UserCount
			global.LOG.Info("用户数量", zap.Int64("UserCount", UserCount))

		}, "定时查看用户数量", option...)
		if err != nil {
			fmt.Println(err.Error())
		}
	}()
}
