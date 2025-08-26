package services

import (
	"time"
)

type CacheService interface {
	GetJSON(key string, dest interface{}) error
	SetJSON(key string, value interface{}, ttl time.Duration) error
}
