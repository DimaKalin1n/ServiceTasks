package serverStart

import (
	"fmt"
	"myApp/internal/handlers"
	mylogger "myApp/internal/myLogger"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func NewServer(logger *mylogger.MyLog) error {
	trace_id, errGen := mylogger.GenerationTraceId()
	if errGen != nil {
		logger.ErrorLog("Не удалось создать trace_id")
	} else {
		logger.InfoLog(trace_id, "удалось создать uuid")
	}

	errEnv := godotenv.Load()
	if errEnv != nil {
		fmt.Println("Ошибка при открытии env file")
	}

	SERVV_PORT := os.Getenv("SERV_PORT")

	server := http.Server{
		Addr:           SERVV_PORT,
		Handler:        handlers.NewHandlerServer(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if errServ := server.ListenAndServe(); errServ != nil {
		logger.ErrorLog("ошибка при запуске сервера")
		return errServ
	}
	logger.InfoLog(trace_id, "Сервер запущен")
	return nil
}
