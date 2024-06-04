package global

import (
	"os"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Viper() *viper.Viper {
	debug := os.Getenv("DEBUG")
	configFileName := "../conf-dev.yaml"
	if debug != "true" {
		configFileName = "../conf-pro.yaml"
	}
	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		zap.S().Fatalf("读取配置文件失败：%s\n", err.Error())
	}
	config = &Config{}
	if err := v.Unmarshal(config); err != nil {
		zap.S().Fatalf("解析配置文件失败：%s\n", err.Error())
	}
	return v
}