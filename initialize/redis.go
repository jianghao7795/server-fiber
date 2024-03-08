package initialize

import (
	"context"

	"server-fiber/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func Redis() {
	// 初始化时可以指定连接redis的读写超时时间，默认都是3s
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.LOG.Error("redis connect ping failed, err:", zap.Error(err))
	} else {
		global.LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.REDIS = client
	}
}
