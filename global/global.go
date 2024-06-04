package global

import (
	"go.uber.org/zap"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	LOG *zap.Logger
	VP *viper.Viper
)