package services

import (
	config_cache "backend/internal/config/cache"
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisCacheService struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisCacheService() *redisCacheService {
	return &redisCacheService{
		client: config_cache.RedisClient,
		ctx:    config_cache.Ctx,
	}
}

func (cs *redisCacheService) SetJSON(key string, value interface{}, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return cs.client.Set(cs.ctx, key, data, ttl).Err()
}

func (cs *redisCacheService) GetJSON(key string, dest interface{}) error {
	val, err := cs.client.Get(cs.ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}
