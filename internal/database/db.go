package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func IninDB() *pgxpool.Pool {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Ошибка получения значений")
		return nil
	}

	var userBD, pasBD, nameBD string = os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("DB_POSTGRES")
	connStr := fmt.Sprintf("postgres://%s:%s@localhost:5432/%s?sslmode=disable", userBD, pasBD, nameBD)

	pool, errPool := pgxpool.New(context.Background(), connStr)
	if errPool != nil {
		fmt.Println("Ошибка в создании пула")
		return nil
	}

	if errPing := pool.Ping(context.Background()); errPing != nil {
		fmt.Println("БД не пингуется")
		return nil
	}

	fmt.Println("Подключение установлено")
	return pool
}
