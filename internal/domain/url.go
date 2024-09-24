package domain

import (
	"context"
	"time"
)

type URL struct {
	ID         string    `db:"id"`
	LongURL    string    `db:"long_url"`
	ShortURL   string    `db:"short_url"`
	VisitCount int       `db:"visit_count"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type URLRepository interface {
	Create(ctx context.Context, longURL string) error
	FindByShortURL(ctx context.Context, shortURL string) error
}

type URLUsecase interface {
}
