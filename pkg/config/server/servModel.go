package server

import (
	"log/slog"
	"net/http"
)

type Server struct {
	Server *http.Server
	log *slog.Logger
}

type ServerMake interface{
	//NewServer(log slog.Logger, port string) (*Server, error)
	StartServer() error
}
