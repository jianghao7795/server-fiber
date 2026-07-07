package middleware

import (
	"server/model/common/response"
	"sync"
	"time"

	"github.com/gofiber/fiber/v3"
)

type TokenBucket struct {
	capacity  int64      // 桶的容量
	rate      float64    // 令牌放入速率
	tokens    float64    // 当前令牌数量
	lastToken time.Time  // 上一次放令牌的时间
	mtx       sync.Mutex // 互斥锁
}

func (tb *TokenBucket) Allow() bool {
	tb.mtx.Lock()
	defer tb.mtx.Unlock()
	now := time.Now()
	// 计算需要放的令牌数量
	tb.tokens = tb.tokens + tb.rate*now.Sub(tb.lastToken).Seconds()
	if max(tb.tokens, float64(tb.capacity)) != tb.tokens {
		tb.tokens = float64(tb.capacity)
	}

	// 判断是否允许请求
	if tb.tokens >= 1 {
		tb.tokens--
		tb.lastToken = now
		return true
	} else {
		return false
	}
}

// defaultBucket 全局共享的令牌桶实例，避免每次请求创建新桶导致限流失效
var defaultBucket = &TokenBucket{
	capacity:  1000,
	rate:      10.0,
	tokens:    0,
	lastToken: time.Now(),
}

// LimitHandler 令牌桶限流中间件，使用全局共享的令牌桶确保限流有效
func LimitHandler(c fiber.Ctx) error {
	if !defaultBucket.Allow() {
		return response.FailWithDetailed(fiber.Map{"msg": "服务器需要休息一下，请等几分钟"}, "加载中", 3, nil, c)
	}
	return c.Next()
}
