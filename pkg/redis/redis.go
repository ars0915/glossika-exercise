package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"

	"github.com/ars0915/glossika-exercise/config"
)

func NewRedisClient(config config.ConfENV) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.Redis.Hosts,
	})

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		return nil, errors.Wrap(err, "redis connect failed")
	}

	return redisClient, nil
}
