package repository

import "url-shortener/internal/pkg/utils/pgxiface"

type ShortenerRepository struct {
	db pgxiface.PgxIface
}

func CreateShortenerRepository(db pgxiface.PgxIface) *ShortenerRepository {
	return &ShortenerRepository{
		db: db,
	}
}

func (r *ShortenerRepository) SaveURL() {
	panic("Not implemented")
}

func (r *ShortenerRepository) GetURL() {
	panic("Not implemented")
}

func (r *ShortenerRepository) DeleteURL() {
	panic("Not implemented")
}
