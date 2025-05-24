package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"go-ecommerce/global"

	"github.com/redis/go-redis/v9"
)

func GetCache(ctx context.Context, key string, obj interface{}) error {
	result, err := global.Redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return fmt.Errorf("key %s not found", key)
	} else if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(result), obj); err != nil {
		return err
	}
	return nil
}
