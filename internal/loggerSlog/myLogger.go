package loggerslog

import (
	"log/slog"
	"os"
)

func CreateLogger(serviceName string) (*slog.Logger, error) {
	logFile, err := os.OpenFile("logFile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	logOpt := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	logger := slog.New(slog.NewJSONHandler(logFile, logOpt))
	slog.SetDefault(logger)

	return logger, nil
}
