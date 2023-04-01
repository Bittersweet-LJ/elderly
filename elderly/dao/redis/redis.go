package redis

import (
	"elderly/settings"
	"fmt"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
)

var (
	client *redis.Client
	//nil    = redis.Nil
)

func Init(cfg *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password,
		DB:       cfg.Db,
		PoolSize: cfg.PoolSize,
	})
	_, err = client.Ping().Result()
	if err != nil {
		zap.L().Error("connect redis failed", zap.Error(err))
		return
	}
	return
}

func Close() {
	_ = client.Close()
}
