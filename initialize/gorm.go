package initialize

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/schema"

	"github.com/congwa/gin-start/global"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBBASE interface {
	GetLogMode() string
}

var Gorm = new(_gorm)

type _gorm struct{}

// Config gorm 自定义配置
// Author [SliverHorn](https://github.com/SliverHorn)
func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {
	Config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,
			SingularTable: singular,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	_default := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: 200 * time.Millisecond,
		LogLevel:      logger.Warn,
		Colorful:      true,
	})
	var logMode DBBASE

	logMode = &global.Config.Mysql

	switch logMode.GetLogMode() {
	case "silent", "Silent":
		Config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		Config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		Config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		Config.Logger = _default.LogMode(logger.Info)
	default:
		Config.Logger = _default.LogMode(logger.Info)
	}

	return Config
}
