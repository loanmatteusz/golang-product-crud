package services

import (
	config_cache "backend/internal/config/cache"
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheService struct {
	client *redis.Client
	ctx    context.Context
}

func NewCacheService() *CacheService {
	return &CacheService{
		client: config_cache.RedisClient,
		ctx:    config_cache.Ctx,
	}
}

func (cs *CacheService) Set(key string, value interface{}, ttl time.Duration) error {
	return cs.client.Set(cs.ctx, key, value, ttl).Err()
}

func (cs *CacheService) Get(key string) (string, error) {
	val, err := cs.client.Get(cs.ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		return "", err
	}
	return val, nil
}

func (cs *CacheService) Delete(key string) error {
	return cs.client.Del(cs.ctx, key).Err()
}

func (cs *CacheService) Exists(key string) (bool, error) {
	count, err := cs.client.Exists(cs.ctx, key).Result()
	return count > 0, err
}

func (cs *CacheService) SetJSON(key string, value interface{}, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return cs.client.Set(cs.ctx, key, data, ttl).Err()
}

func (cs *CacheService) GetJSON(key string, dest interface{}) error {
	val, err := cs.client.Get(cs.ctx, key).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}
