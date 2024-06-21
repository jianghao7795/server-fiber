package initialize

import (
	"context"

	"server-fiber/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func Redis() error {
	// 初始化时可以指定连接redis的读写超时时间，默认都是3s
	redisCfg := global.CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,     // redis服务ip:port
		Password: redisCfg.Password, // redis的认证密码
		DB:       redisCfg.DB,       // 连接的database
		//IdleTimeout: 300,               // 默认Idle超时时间
		PoolSize: 100, // 连接池
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.LOG.Error("redis connect ping failed, err:", zap.Error(err))
		return err
	} else {
		global.LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.REDIS = client
	}
	return err
}
