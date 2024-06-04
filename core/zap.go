package core

import (
	"go.uber.org/zap"
)
func Zap() *zap.Logger {
	// 初始化logger
	logger, _ := zap.NewDevelopment()
	// 在代码的任何地方都可以直接使用 zap.Info(), zap.Error() 等全局函数记录日志,而不需要显式地传入日志记录器实例
	zap.ReplaceGlobals(logger)
	
	return logger
}