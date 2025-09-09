package serverStart

import (
	"fmt"
	"log/slog"
	"myApp/internal/handlers"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func NewServer(log *slog.Logger) error {
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
		log.Error("ошибка при запуске сервера", "error:", errServ.Error())
		return errServ
	}
	log.Info("Сервер запущен")
	return nil
}
