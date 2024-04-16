package global

import (
	"log/slog"
	"server-fiber/config"
	"server-fiber/utils/timer"

	ut "github.com/go-playground/universal-translator"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB // gorm
	DBList map[string]*gorm.DB
	REDIS  *redis.Client // redis
	CONFIG config.Server // 配置
	VIP    *viper.Viper  // 读取配置文件
	// LOG    *oplogging.Logger
	LOG                *zap.Logger // 日志
	Timer              = timer.NewTimerTask()
	ConcurrencyControl = &singleflight.Group{} // 记录token

	BlackCache local_cache.Cache // 缓存
	// lock       sync.RWMutex
	Validate ut.Translator
	// 缓存
	// Cache config.Cache
	Logger *slog.Logger // 用处 打印log
)

func Done(c *fiber.Ctx, logString []byte) {
	if c.Response().StatusCode() >= fiber.StatusBadRequest {
		if c.Response().StatusCode() == 404 {
			LOG.Error(string(logString))
		} else {
			LOG.Warn(string(logString))
		}
	}
}

// GetGlobalDBByDBName 通过名称获取db list中的db
// func GetGlobalDBByDBName(dbname string) *gorm.DB {
// 	lock.RLock()
// 	defer lock.RUnlock()
// 	return DBList[dbname]
// }

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
// func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
// 	lock.RLock()
// 	defer lock.RUnlock()
// 	db, ok := DBList[dbname]
// 	if !ok || db == nil {
// 		panic("db no init")
// 	}
// 	return db
// }
