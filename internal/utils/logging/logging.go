package logging

import "github.com/sirupsen/logrus"

var Logger = logrus.New()

func SetupLogger() {
	// установка уровня логгирования
	Logger.SetLevel(logrus.InfoLevel) // Debug не логгируется

	// логи цветные и выводится время полностью
	textFormatter := &logrus.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	}

	// установка формата вывода логов
	Logger.SetFormatter(textFormatter)
}
