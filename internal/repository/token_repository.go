package repository

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type TokenRepository struct {
	Redis *redis.Client
}

func NewTokenRepository(redis *redis.Client) *TokenRepository {
	return &TokenRepository{
		Redis: redis,
	}
}

func (r *TokenRepository) SetToken(token string) error {
	return r.Redis.Set(
		context.Background(),
		"refresh:"+token,
		"1",
		24*time.Hour,
	).Err()
}

func (r *TokenRepository) GetToken(token string) (string, error) {
	id, err := r.Redis.Get(
		context.Background(),
		"refresh:"+token,
	).Result()

	if err != nil {
		return "", errors.New("refresh token tidak valid")
	}

	return id, nil
}

func (r *TokenRepository) DeleteToken(token string) error {
	return r.Redis.Del(
		context.Background(),
		"refresh:"+token,
	).Err()
}
