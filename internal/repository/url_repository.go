package repository

import (
	"context"
	"database/sql"
	"simple-go-shorten/internal/domain"
)

type urlRepository struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) domain.URLRepository {
	return &urlRepository{db: db}
}

// Create implements domain.URLRepository.
func (r *urlRepository) Create(ctx context.Context, longURL string) error {
	panic("unimplemented")
}

// FindByShortURL implements domain.URLRepository.
func (r *urlRepository) FindByShortURL(ctx context.Context, shortURL string) error {
	panic("unimplemented")
}
