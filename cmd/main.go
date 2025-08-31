package main

import (
	"fmt"
	"myApp/internal/database"
	"myApp/internal/server"
	tasktype "myApp/internal/taskType"
	"myApp/internal/user"
	"net/http"
)

func main() {

	dbPool := database.IninDB()
	defer dbPool.Close()
	srv := &server.Server{DB: dbPool}

	http.HandleFunc("/login", user.Login(srv))
	http.HandleFunc("/createUser", user.Register(srv))
	http.HandleFunc("/createTaskType", tasktype.CreateTaskType(srv))
	http.HandleFunc("/getAllType", tasktype.GetAllTasksType(srv))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Ошибка сервера")
	}
}
