package repository

import (
	"context"
	"url-shortener/internal/models"
	"url-shortener/internal/pkg/utils/logging"
	"url-shortener/internal/pkg/utils/pgxiface"
)

type ShortenerRepository struct {
	db pgxiface.PgxIface
}

func CreateShortenerRepository(db pgxiface.PgxIface) *ShortenerRepository {
	return &ShortenerRepository{
		db: db,
	}
}

func (r *ShortenerRepository) SaveURL(ctx context.Context, urlToSave string, alias string) (newURLStruct models.URLStruct, err error) {
	funcName := "SaveURL"
	query := `INSERT INTO url (url, alias)
			VALUES ($1, $2) RETURNING id, url, alias;`

	err = r.db.QueryRow(ctx, query, urlToSave, alias).Scan(
		&newURLStruct.ID,
		&newURLStruct.URL,
		&newURLStruct.Alias,
	)

	logging.Logger.Debugf("%s query has err: %v", funcName, err)

	return newURLStruct, nil
}

func (r *ShortenerRepository) GetURL() {
	panic("Not implemented")
}

func (r *ShortenerRepository) DeleteURL() {
	panic("Not implemented")
}
