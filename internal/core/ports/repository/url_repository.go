package repository

import (
	"context"

	"github.com/not201/ninja-url/internal/core/domain/entity"
)

type UrlRepository interface {
	Save(ctx context.Context, url *entity.Url) error
	FindByShortCode(ctx context.Context, shortCode string) (string, error)
}
