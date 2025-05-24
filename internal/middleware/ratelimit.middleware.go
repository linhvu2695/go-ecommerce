package middleware

import (
	"fmt"
	"go-ecommerce/global"
	"go-ecommerce/pkg/utils/context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	redisStore "github.com/ulule/limiter/v3/drivers/store/redis"
	"go.uber.org/zap"
)

type RateLimiter struct {
	globalRateLimiter  *limiter.Limiter
	publicRateLimiter  *limiter.Limiter
	privateRateLimiter *limiter.Limiter
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		globalRateLimiter:  rateLimiter("100-S"),
		publicRateLimiter:  rateLimiter("80-S"),
		privateRateLimiter: rateLimiter("50-S"),
	}
}

func rateLimiter(interval string) *limiter.Limiter {
	store, err := redisStore.NewStoreWithOptions(global.Redis, limiter.StoreOptions{
		Prefix:          "rate-limiter",
		MaxRetry:        3,
		CleanUpInterval: time.Hour,
	})
	if err != nil {
		return nil
	}

	rate, err := limiter.NewRateFromFormatted(interval)
	if err != nil {
		panic(err)
	}

	instance := limiter.New(store, rate)
	return instance
}

func (rl *RateLimiter) FilterLimitUrlPath(url string) *limiter.Limiter {
	if url == "/v1/user/login" {
		return rl.publicRateLimiter
	} else if url == "v1/user/info" {
		return rl.privateRateLimiter
	} else {
		return rl.globalRateLimiter
	}
}

func (rl *RateLimiter) GlobalRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := "GLOBAL"

		limitContext, err := rl.globalRateLimiter.Get(c, key)
		if err != nil {
			global.Logger.Error(fmt.Sprintf("Failed to check rate limit %s", key), zap.Error(err))
			c.Next()
			return
		}

		if limitContext.Reached {
			global.Logger.Warn(fmt.Sprintf("Rate limit breached %s limit", key))
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": fmt.Sprintf("Rate limit breached %s, try later", key)})
			return
		}

		c.Next()
	}
}

func (rl *RateLimiter) PublicRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path
		limiter := rl.FilterLimitUrlPath(urlPath)

		if limiter != nil {
			key := fmt.Sprintf("%s-%s", c.ClientIP(), urlPath)

			limitContext, err := rl.globalRateLimiter.Get(c, key)
			if err != nil {
				global.Logger.Error(fmt.Sprintf("Failed to check rate limit %s", key), zap.Error(err))
				c.Next()
				return
			}

			if limitContext.Reached {
				global.Logger.Warn(fmt.Sprintf("Rate limit breached %s limit", key))
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": fmt.Sprintf("Rate limit breached %s, try later", key)})
				return
			}
		}

		c.Next()
	}
}

func (rl *RateLimiter) PrivateRateLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		urlPath := c.Request.URL.Path
		uuid := c.ClientIP()
		userId, err := context.GetUserID(c.Request.Context())
		if err == nil {
			uuid = string(userId)
		}

		limiter := rl.FilterLimitUrlPath(urlPath)

		if limiter != nil {
			key := fmt.Sprintf("%s-%s", uuid, urlPath)

			limitContext, err := rl.globalRateLimiter.Get(c, key)
			if err != nil {
				global.Logger.Error(fmt.Sprintf("Failed to check rate limit %s", key), zap.Error(err))
				c.Next()
				return
			}

			if limitContext.Reached {
				global.Logger.Warn(fmt.Sprintf("Rate limit breached %s limit", key))
				c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": fmt.Sprintf("Rate limit breached %s, try later", key)})
				return
			}
		}

		c.Next()
	}
}
