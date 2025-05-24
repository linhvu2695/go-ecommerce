package initialization

import (
	"context"
	"fmt"
	"go-ecommerce/global"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func InitRedis() {
	r := global.Config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database, // use default DB
		PoolSize: 10,
	})

	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Failed to initialize Redis", zap.Error(err))
	}

	global.Logger.Info("Initialize Redis successfully")
	global.Redis = rdb

	redisExample()
}

func redisExample() {
	ctx := context.Background()
	err := global.Redis.Set(ctx, "score", 100, 0).Err()
	if err != nil {
		fmt.Println("Error redis example:", zap.Error(err))
		return
	}

	value, err := global.Redis.Get(ctx, "score").Result()
	if err != nil {
		fmt.Println("Error redis example:", zap.Error(err))
		return
	}

	global.Logger.Info("Redis example:", zap.String("score", value))
}
