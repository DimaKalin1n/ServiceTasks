package user

import (
	"context"
	"encoding/json"
	"fmt"
	"myApp/internal/auth"
	"myApp/internal/server"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Register(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Accept", "*/*")
		if r.Method != http.MethodPost {
			w.WriteHeader(405)
			str := string(r.Method) + "rule of" + r.Pattern + " не используется"
			w.Write([]byte(str))
			return
		}
		var UserInfo struct {
			Login    string `json:"email"`
			Password string `json:"password"`
			TimeCr   time.Time
		}
		if err := json.NewDecoder(r.Body).Decode(&UserInfo); err != nil {
			w.WriteHeader(503)
			w.Write([]byte("Ошибка сервера не удалось распарсить json"))
			return
		}

		if len(UserInfo.Login) < 3 || len(UserInfo.Password) < 8 {
			w.WriteHeader(400)
			w.Write([]byte("Невалидные данные"))
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(UserInfo.Password), bcrypt.DefaultCost)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Ошибка при хэшировании пароля"))
			return
		}

		query := `INSERT INTO users (email, password) VALUES ($1, $2)`
		_, err = s.DB.Exec(context.Background(), query, UserInfo.Login, string(hashedPassword))
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Ошибка при добавлении пользователя в БД: " + err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Пользователь успешно зарегистрирован",
		})
	}
}

func Login(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(409)
			w.Write([]byte(string(r.Method) + "rule of" + r.Pattern + " не используется"))
			return
		}

		var User struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}

		err := json.NewDecoder(r.Body).Decode(&User)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Ошибка сервера"))
			return
		}

		var UserInfo struct {
			id       int
			email    string
			password string
		}
		query := `SELECT id, email, password FROM users WHERE email = $1`
		errSelect := s.DB.QueryRow(context.Background(), query, User.Email).Scan(&UserInfo.id, &UserInfo.email, &UserInfo.password)
		fmt.Println(UserInfo)
		if errSelect != nil {
			w.WriteHeader(503)
			w.Write([]byte("Ошибка сервера"))
			return
		}

		errComp := bcrypt.CompareHashAndPassword([]byte(UserInfo.password), []byte(User.Password))
		if errComp != nil {
			w.WriteHeader(403)
			w.Write([]byte("Неверный логин или пароль"))
			return
		}

		tok, err := auth.GenerateToken(UserInfo.id, UserInfo.email)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(401)
			w.Write([]byte("Не удалось создать токен"))
			return
		}
		fmt.Println(tok)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Успешный вход",
		})

	}
}
