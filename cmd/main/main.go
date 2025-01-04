package main

import (
	"fmt"
	"os"
	"url-shortener/internal/pkg/config"
	"url-shortener/internal/pkg/utils/db"
	"url-shortener/internal/pkg/utils/logging"

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

	// TODO: роутер

	// TODO: run server
}
