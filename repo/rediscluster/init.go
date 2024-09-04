package rediscluster

import (
	"github.com/go-redis/redis/v8"
)

type RedisRepo struct {
	client *redis.ClusterClient
}

func New(client *redis.ClusterClient) *RedisRepo {
	return &RedisRepo{
		client: client,
	}
}

func (r RedisRepo) Close() error {
	return r.client.Close()
}
