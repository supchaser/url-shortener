package main

import (
	"fmt"
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/utils/logging"

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

	// TODO: база данных

	// TODO: роутер

	// TODO: run server
}
