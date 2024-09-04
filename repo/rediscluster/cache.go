package rediscluster

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

func (r RedisRepo) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	if err := r.client.Set(ctx, key, value, expiration).Err(); err != nil {
		return errors.Wrap(err, "error setting value")
	}
	return nil
}

func (r RedisRepo) Get(ctx context.Context, key string) (interface{}, error) {
	return r.client.Get(ctx, key).Result()
}
