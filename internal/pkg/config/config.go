package config

import (
	"os"
	"time"
	"url-shortener/internal/pkg/utils/logging"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	DBURL      string `yaml:"db_url" env-required:"true"`
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-defalult:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeput time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

var (
	CurrentConfig *Config
)

// TODO: прокидывать подключение к БД через ENV
func MustLoad() *Config {
	configPath := os.Getenv("PATH_CONFIG")
	if configPath == "" {
		logging.Logger.Fatal("PATH_CONFIG is not set")
	}

	// dbURL := os.Getenv("DB_URL")
	// if dbURL == "" {
	// 	logging.Logger.Fatal("DB_URL is not set")
	// }

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		logging.Logger.WithField("configPath", configPath).Fatal("config file does not exist")
	}

	CurrentConfig = &Config{}
	if err := cleanenv.ReadConfig(configPath, CurrentConfig); err != nil {
		logging.Logger.WithField("error", err).Fatal("cannot read file")
	}

	return CurrentConfig
}
