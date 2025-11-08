package repositories

import (
	"context"
	"time"

	"github.com/not201/ninja-url/internal/core/domain/entity"
	"github.com/not201/ninja-url/internal/core/ports/repository"
	"github.com/redis/go-redis/v9"
)

type urlRepository struct {
	redis *redis.Client
}

func NewUrlRepository(redis *redis.Client) repository.UrlRepository {
	return &urlRepository{redis: redis}
}

func (r *urlRepository) Save(ctx context.Context, url *entity.Url) error {
	key := url.ShortCode
	value := url.OriginalUrl
	ttl := time.Until(url.ExpiresAt)

	return r.redis.Set(ctx, key, value, ttl).Err()
}

func (r *urlRepository) FindByShortCode(ctx context.Context, shortCode string) (string, error) {
	key := shortCode
	value, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return value, nil
}
