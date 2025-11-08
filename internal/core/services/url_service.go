package services

import (
	"context"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/not201/ninja-url/internal/core/domain/dto"
	"github.com/not201/ninja-url/internal/core/domain/entity"
	"github.com/not201/ninja-url/internal/core/ports/repository"
	"github.com/not201/ninja-url/internal/core/ports/service"
)

type urlService struct {
	repo repository.UrlRepository
}

func NewUrlService(repo repository.UrlRepository) service.UrlService {
	return &urlService{repo: repo}
}

func (s *urlService) Shorten(ctx context.Context, dto *dto.UrlDto) (*entity.Url, error) {

	now := time.Now()

	shortCode, err := gonanoid.New(6)
	if err != nil {
		return nil, err
	}

	url := entity.Url{
		OriginalUrl: dto.OriginalUrl,
		ShortCode:   shortCode,
		ExpiresAt:   now.Add(time.Hour * 24),
	}

	if err := s.repo.Save(ctx, &url); err != nil {
		return nil, err
	}

	return &url, nil
}

func (s *urlService) Resolver(ctx context.Context, shortCode string) (string, error) {
	originalUrl, err := s.repo.FindByShortCode(ctx, shortCode)
	if err != nil {
		return "", err
	}

	return originalUrl, nil
}
