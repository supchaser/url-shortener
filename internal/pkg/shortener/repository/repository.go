package repository

import (
	"context"
	"errors"
	"fmt"
	"url-shortener/internal/models"
	"url-shortener/internal/pkg/utils/logging"
	"url-shortener/internal/pkg/utils/pgxiface"

	"github.com/jackc/pgx/v5"
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

	return newURLStruct, err
}

func (r *ShortenerRepository) GetURL(ctx context.Context, alias string) (*models.URLStruct, error) {
	funcName := "GetURL"
	query := `SELECT id, url, alias FROM url WHERE alias=$1;`
	row := r.db.QueryRow(ctx, query, alias)
	url := &models.URLStruct{}
	err := row.Scan(
		&url.ID,
		&url.URL,
		&url.Alias,
	)

	logging.Logger.Debugf("%s with alias='%s' query has error: %v", funcName, alias, err)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("%s: no rows found: %w", funcName, err)
		}
		return nil, fmt.Errorf("%s: scanning error: %w", funcName, err)
	}

	return url, nil
}

func (r *ShortenerRepository) DeleteURL(ctx context.Context, alias string) (err error) {
	funcName := "DeleteURL"
	if alias == "" {
		return fmt.Errorf("%s: alias is empty", funcName)
	}

	query := `DELETE FROM url WHERE alias=$1;`
	url, err := r.db.Exec(ctx, query, alias)
	if url.RowsAffected() == 0 {
		return fmt.Errorf("%s: not found", funcName)
	}

	logging.Logger.Debugf("%s query has error: %v", funcName, err)
	if err != nil {
		return err
	}

	return nil
}
