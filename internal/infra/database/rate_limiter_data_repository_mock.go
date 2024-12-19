package database

import (
	"context"

	"github.com/flaviotsx/rate-limiter/internal/entity"
	"github.com/stretchr/testify/mock"
)

type RateLimiterDataRepositoryMock struct {
	mock.Mock
}

func (m *RateLimiterDataRepositoryMock) Find(ctx context.Context, key string) (*entity.RateLimiterData, error) {
	args := m.Called(ctx, key)
	return args.Get(0).(*entity.RateLimiterData), args.Error(1)
}

func (m *RateLimiterDataRepositoryMock) Save(ctx context.Context, key string, data *entity.RateLimiterData) error {
	args := m.Called(ctx, key, data)
	return args.Error(0)
}
