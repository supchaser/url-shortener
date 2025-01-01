package main

import "url-shortener/internal/utils/logging"

func main() {
	// TODO: конфиг

	// TODO: логгер
	logger := logging.SetupLogger()
	logger.Info("The logger has been successfully initialized")

	// TODO: база данных

	// TODO: роутер

	// TODO: run server
}
