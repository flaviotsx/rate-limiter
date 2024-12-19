package database

import (
	"context"
	"encoding/json"
	"time"

	"github.com/flaviotsx/rate-limiter/internal/entity"
	"github.com/redis/go-redis/v9"
)

type RedisRateLimiterDataRepository struct {
	db *redis.Client
}

func NewRedisRateLimiterDataRepository(db *redis.Client) *RedisRateLimiterDataRepository {
	return &RedisRateLimiterDataRepository{
		db: db,
	}
}

func (r *RedisRateLimiterDataRepository) Find(ctx context.Context, key string) (*entity.RateLimiterData, error) {
	found, err := r.db.Get(ctx, key).Result()

	if err == redis.Nil {
		return &entity.RateLimiterData{
			Requests: 0,
			LastSeen: time.Now(),
			Blocked:  false,
		}, nil
	} else if err != nil {
		return nil, err
	}

	var data *entity.RateLimiterData
	err = json.Unmarshal([]byte(found), &data)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *RedisRateLimiterDataRepository) Save(ctx context.Context, key string, data *entity.RateLimiterData) error {
	dataEncoded, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return r.db.Set(ctx, key, dataEncoded, 0).Err()
}
