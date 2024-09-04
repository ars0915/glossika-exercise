package rediscluster

import (
	"github.com/go-redis/redis/v8"
)

type RedisRepo struct {
	client *redis.Client
}

func New(client *redis.Client) *RedisRepo {
	return &RedisRepo{
		client: client,
	}
}

func (r RedisRepo) Close() error {
	return r.client.Close()
}
