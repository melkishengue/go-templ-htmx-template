package logger

import (
	"log"
	"log/slog"
	"os"
)

func SetUp(env string) {
	logLovel := new(slog.LevelVar)
	logLovel.Set(slog.LevelInfo)
	if env == "development" {
		logLovel.Set(slog.LevelDebug)
	}

	_logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(_logger)
	slog.SetLogLoggerLevel(slog.LevelDebug)

	log.Println("Log level is set to", logLovel.Level())
}
