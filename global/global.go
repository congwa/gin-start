package global

import (
	"github.com/congwa/gin-start/config"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 定义全局变量
var (
	DB     *gorm.DB
	LOG    *zap.Logger
	VP     *viper.Viper
	Config config.Config
	Redis  redis.UniversalClient
)
