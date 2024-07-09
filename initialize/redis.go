package initialize

import (
	"context"

	"github.com/congwa/gin-start/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := global.Config.Redis
	var client redis.UniversalClient
	client = redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.LOG.Error("redis connect ping failed, err:", zap.Error(err))
		panic(err)
	} else {
		global.LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.Redis = client
	}
}
