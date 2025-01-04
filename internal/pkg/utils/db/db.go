package db

import (
	"context"
	"fmt"
	"time"
	"url-shortener/internal/pkg/config"
	"url-shortener/internal/pkg/utils/logging"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnnectToPgx() (db *pgxpool.Pool, err error) {
	config, err := pgxpool.ParseConfig(config.CurrentConfig.DBURL)
	if err != nil {
		return nil, fmt.Errorf("error creating pgx config: %w", err)
	}

	db, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("error creating PostgreSQL connection: %w", err)
	}

	for range 10 {
		if err = db.Ping(context.Background()); err == nil {
			return db, nil
		}
		logging.Logger.Warn("Retry Postgres ping")
		time.Sleep(1 * time.Second)
	}

	return nil, fmt.Errorf("error while pinging PostgreSQL: %w", err)
}
