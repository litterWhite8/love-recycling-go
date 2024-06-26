package bootstrap

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"love-recycling-go/global"
)

func InitializeRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: global.App.Config.Redis.Host + ":" + global.App.Config.Redis.Port,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.App.Log.Error("Redis connect ping failed, err:", zap.Any("err", err))
		return nil
	}
	return client
}
