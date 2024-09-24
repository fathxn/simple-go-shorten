package usecase

import "simple-go-shorten/internal/domain"

type urlUsecase struct {
	URLRepository domain.URLRepository
}

func NewURLUsecase(urlRepository domain.URLRepository) domain.URLUsecase {
	return &urlUsecase{URLRepository: urlRepository}
}
