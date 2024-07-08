package main

import (
	"fmt"

	"github.com/congwa/gin-start/initialize"
	"github.com/joho/godotenv"

	// 全局模块

	"github.com/congwa/gin-start/core"
	"github.com/congwa/gin-start/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db = make(map[string]string)

// 初始化数据库
// 使用gorm
func initDB() *gorm.DB {
	m := global.Config.Mysql
	mysqlConfig := mysql.Config{
		DSN:                       m.Dsn(), // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}

	if db, err := gorm.Open(mysql.New(mysqlConfig), initialize.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

func main() {
	// 增加 .env 的能力
	err := godotenv.Load()
	if err != nil {
		// 如果 .env 文件不存在,则使用默认值
		// 这里只是一个示例,实际应用中应该有更健壮的错误处理
		fmt.Println("No .env file found")
	}

	// 从core中初始化 配置 模块
	global.VP = core.Viper()
	// 从core中初始化 日志 模块
	global.LOG = core.Zap()

	global.DB = initDB()

	core.RunServer()
}
