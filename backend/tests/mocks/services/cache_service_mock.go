package mocks

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type CacheService struct {
	mock.Mock
}

func (m *CacheService) GetJSON(key string, dest interface{}) error {
	args := m.Called(key, dest)
	return args.Error(0)
}

func (m *CacheService) SetJSON(key string, value interface{}, ttl time.Duration) error {
	args := m.Called(key, value, ttl)
	return args.Error(0)
}
