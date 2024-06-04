package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	// 全局模块
	"github.com/congwa/gin-start/global" 
	"github.com/congwa/gin-start/config"
	"github.com/congwa/gin-start/core"
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
func initDB() {
	var err error
}

func main() {
	// 增加 .env 的能力
	err := godotenv.Load()
	if err != nil {
		// 如果 .env 文件不存在,则使用默认值
		// 这里只是一个示例,实际应用中应该有更健壮的错误处理
		fmt.Println("No .env file found")
	}

	global.VP = core.Viper()
	// 从core中初始化 日志 模块
	global.LOG = core.Zap()



	// 获取配置 配置这里使用 viper
	conf := config.GetConfig()



	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
