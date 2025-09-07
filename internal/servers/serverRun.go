package servers

import (
	"myApp/internal/loggerSlog"
	"myApp/routes"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)




func NewServer(sl slog.Logger) error{
	if err  := godotenv.Load(); err != nil{
		sl.Error("Ошибка при получении env файла " + err.Error())
		return err
	}

	port := os.Getenv("SERV_PORT")

	ser := &http.Server{
	Addr: ":" + port,
	Handler: routes.CreateMux(),
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
}
	if err := ser.ListenAndServe(); err != nil{
		sl.Error("Ошибка при запуске сервера " + err.Error())
		return err
	}
}