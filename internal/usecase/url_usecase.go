package usecase

import (
	"context"
	"simple-go-shorten/internal/domain"
)

type urlUsecase struct {
	URLRepository domain.URLRepository
}

func NewURLUsecase(urlRepository domain.URLRepository) domain.URLUsecase {
	return &urlUsecase{URLRepository: urlRepository}
}

func (u *urlUsecase) Create(ctx context.Context, longURL string) error {
	return nil
}
func (u *urlUsecase) Redirect(ctx context.Context) error {
	return nil
}
