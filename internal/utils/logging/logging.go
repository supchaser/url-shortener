package logging

import "github.com/sirupsen/logrus"

func SetupLogger() *logrus.Logger {
	logger := logrus.New()

	// установка уровня логгирования
	logger.SetLevel(logrus.InfoLevel) // Debug не логгируется

	// логи цветные и выводится время полностью
	textFormatter := &logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	}

	// установка формата вывода логов
	logger.SetFormatter(textFormatter)

	return logger
}
