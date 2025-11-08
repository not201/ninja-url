package service

import (
	"context"

	"github.com/not201/ninja-url/internal/core/domain/dto"
	"github.com/not201/ninja-url/internal/core/domain/entity"
)

type UrlService interface {
	Shorten(ctx context.Context, dto *dto.UrlDto) (*entity.Url, error)
	Resolver(ctx context.Context, shortCode string) (string, error)
}
