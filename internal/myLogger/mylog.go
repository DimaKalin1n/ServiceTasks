package mylogger

import (
	"log/slog"
	"os"
)

func NewMyLogger() *slog.Logger {

	file, err := os.OpenFile("internal/myLogger/logFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		slog.Error("Не удалось открыть файл для записи логов")
		return nil
	}

	logOpt := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	logHandler := slog.NewJSONHandler(file, logOpt)

	mylog := slog.New(logHandler).With("service", "myApp")

	return mylog
}
