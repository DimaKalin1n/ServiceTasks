package tasktype

import (
	"context"
	"encoding/json"
	"myApp/internal/server"
	"net/http"
)

type TaskType struct {
	Code       string `json:"code"`
	NameUser   string `json:"nameUser"`
	NameClient string `json:"nameClient"`
	active     bool   `json:"active,omitempty"`
}

func CreateTaskType(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(409)
			w.Write([]byte("Метод " + r.Method + " не поддурживается по данному адресу"))
			return
		}
		var taskType TaskType

		err := json.NewDecoder(r.Body).Decode(&taskType)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Не удалось получить данные"))
			return
		}

		query := `INSERT INTO tasktype (code, nameuser, nameclient, active) VALUES ($1, $2, $3, $4)`

		_, errIns := s.DB.Exec(context.Background(), query, taskType.Code, taskType.NameUser, taskType.NameClient, taskType.active)
		if errIns != nil {
			w.WriteHeader(503)
			w.Write([]byte("Не удалось записать значние, ошибка"))
			return
		}

		w.WriteHeader(201)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Тип таска создан",
		})
	}
}

func GetAllTasksType(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(409)
			w.Write([]byte("Метод " + r.Method + " не поддурживается по данному адресу"))
			return
		}
		var typeList []TaskType
		query := `SELECT * FROM tasktype`
		rows, err := s.DB.Query(context.Background(), query)
		defer rows.Close()

		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Ошибка запроса к серверу"))
			return
		}

		for rows.Next() {
			var taskType TaskType
			errPars := rows.Scan(&taskType.Code, &taskType.NameUser, &taskType.NameClient, &taskType.active)
			if errPars != nil {
				w.WriteHeader(503)
				w.Write([]byte("Ошибка парсинга запроса"))
				return
			}
			typeList = append(typeList, taskType)
		}
		_, errMarsh := json.Marshal(typeList)
		if errMarsh != nil {
			w.WriteHeader(503)
			w.Write([]byte("Ошибка преобразования json для ответа"))
			return
		}
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(typeList)

	}
}
