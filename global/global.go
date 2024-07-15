package global

import (
	"server-fiber/config"
	"server-fiber/utils/timer"
	"sync"

	ut "github.com/go-playground/universal-translator"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

// Cache config.Cache 使用 redis

var (
	DB                 *gorm.DB // gorm
	DBList             map[string]*gorm.DB
	REDIS              *redis.Client // redis
	CONFIG             config.Server // 配置
	VIP                *viper.Viper  // 读取配置文件
	LOG                *zap.Logger   // 日志 打印日志 debug fatal error info warn 等几种方式
	Timer              = timer.NewTimerTask()
	ConcurrencyControl = &singleflight.Group{} // 记录token

	BlackCache local_cache.Cache // 缓存
	// Validate lock       sync.RWMutex
	Validate ut.Translator
	//Logger   *slog.Logger // 用处 打印log
	lock sync.RWMutex
)

func Done(c *fiber.Ctx, logString []byte) {
	if c.Response().StatusCode() >= fiber.StatusBadRequest {
		if c.Response().StatusCode() == 404 {
			LOG.Info(string(logString))
		} else {
			LOG.Warn(string(logString))
		}
	}
}

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
