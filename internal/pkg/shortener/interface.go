package shortener

import (
	"context"
	"url-shortener/internal/models"
)

type ShortenerRepository interface {
	SaveURL(ctx context.Context, urlToSave string, alias string) (*models.URLStruct, error)
	GetURL(ctx context.Context, alias string) (*models.URLStruct, error)
	DeleteURL(ctx context.Context, alias string) (err error)
}
