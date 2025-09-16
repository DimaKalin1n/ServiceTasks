package server

import (
	"log/slog"
	"ServiceTasks/internal/handlers"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func NewServer(log *slog.Logger, port string) (*Server, error) {

	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Info("Ошибка при открытии env file")
		return nil, errEnv
	}

	SERVV_PORT := os.Getenv(port)
	if SERVV_PORT == "" {
		log.Error("Не обнаружено значение в файле env, сервер будет запущен на 8080")
		SERVV_PORT = "8080"
	}

	serverHttp := http.Server{
		Addr:           ":" + SERVV_PORT,
		Handler:        handlers.NewHandlerServer(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		IdleTimeout:    120 * time.Second,
	}

	server := &Server{
		Server: &serverHttp,
		log:    log,
	}

	log.Info("Server создан")
	return server, nil
}

func (s *Server) StartServer() error {
	s.log.Info("Сервер запущен, порт " + s.Server.Addr)
	return s.Server.ListenAndServe()
}
