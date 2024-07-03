package main

import (
	"fmt"
	"net/http"

	"github.com/congwa/gin-start/initialize"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// 全局模块

	"github.com/congwa/gin-start/core"
	"github.com/congwa/gin-start/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

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

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
