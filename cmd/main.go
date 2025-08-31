package main

import (
	"context"
	"fmt"
	"log"
	"myApp/internal/server"
	"myApp/internal/user"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v5"
)

func main() {
	if err := godotenv.Load(); err != nil{
		fmt.Print("ошибка загрузки env")
	}
	var  bdUser, bdUsPas, bdName string =  
	os.Getenv("POSTGRES_USER"), 
	os.Getenv("POSTGRES_PASSWORD"), 
	os.Getenv("DB_POSTGRES")
	
	
	
	connStr := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", bdUser,bdUsPas,bdName )

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		fmt.Println("не удалось подключиться к БД")
	}
	defer conn.Close(context.Background())
	if err := conn.Ping(context.Background()); err != nil {
		log.Fatalf("БД недоступна: %v", err)
	}
	srv := &server.Server{DB: conn}

	http.HandleFunc("/login", user.Login(srv))
	http.HandleFunc("/createUser", user.Register(srv))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка сервера")
	}
}
