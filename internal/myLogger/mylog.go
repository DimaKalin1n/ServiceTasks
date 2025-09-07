package mylogger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gofrs/uuid"
)

type MyLog struct {
	slog *slog.Logger
}

func GenerationTraceId() (string, error) {
	randomUuid, err := uuid.NewV4()
	if err != nil {
		fmt.Println("ошибка создания uuid " + err.Error())
		return "", err
	}
	return randomUuid.String(), nil
}

func NewMyLogger() *MyLog {

	file, err := os.OpenFile("logFile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("Не удалось открыть файл для работы")
		return nil
	}

	logOpt := &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}

	logHandler := slog.NewJSONHandler(file, logOpt)

	mylog := slog.New(logHandler).With("service", "myApp")

	return &MyLog{slog: mylog}
}

func (log *MyLog) InfoLog(request_id string, msg string) {
	log.slog.Info(msg)
}

func (log *MyLog) ErrorLog(msg string) {
	log.slog.Error(msg)
}
