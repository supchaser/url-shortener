package shortener

import (
	"context"
	"url-shortener/internal/models"
)

type ShortenerRepository interface {
	SaveURL(ctx context.Context, urlToSave string, alias string) (newURLStruct models.URLStruct, err error)
}
