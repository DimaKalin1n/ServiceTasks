package server

import "github.com/jackc/pgx/v5"

type Server struct {
	DB *pgx.Conn
}
