package main

import (
	"context"
	"fmt"
	"os"
	"url-shortener/internal/pkg/config"
	mlog "url-shortener/internal/pkg/middleware/logging"
	"url-shortener/internal/pkg/shortener/repository"
	"url-shortener/internal/pkg/utils/db"
	"url-shortener/internal/pkg/utils/logging"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		fmt.Println("Ошибка загрузки .env файла:", err)
		os.Exit(1)
	}

	// конфиг
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// логгер
	logging.SetupLogger()
	logging.Logger.Info("The logger has been successfully initialized")

	postgresDB, err := db.ConnnectToPgx()
	if err != nil {
		logging.Logger.Error("error connecting to PostgreSQL: ", err)
		return
	}
	defer postgresDB.Close()

	repoShortener := repository.CreateShortenerRepository(postgresDB)
	urlStruct, err := repoShortener.SaveURL(context.Background(), "https://google.com", "google")
	if err != nil {
		logging.Logger.Error("error saving URL: ", err)
		return
	}

	fmt.Printf("URL has been saved: %+v\n", urlStruct)

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	// router.Use(middleware.Logger) можно как альтернатива самописному
	router.Use(mlog.NewLogrusLogger())
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	// TODO: run server
}
