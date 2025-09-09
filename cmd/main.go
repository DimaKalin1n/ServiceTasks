package main

import (
	"myApp/internal/database"
	mylogger "myApp/internal/myLogger"
	"myApp/internal/server"
	tasktype "myApp/internal/taskType"
	"myApp/internal/user"
	"net/http"
	"myApp/internal/serveStart"
)

func main() {

	dbPool := database.IninDB()
	defer dbPool.Close()
	srv := &server.Server{DB: dbPool}

	http.HandleFunc("/login", user.Login(srv))
	http.HandleFunc("/createUser", user.Register(srv))
	http.HandleFunc("/createTaskType", tasktype.CreateTaskType(srv))
	http.HandleFunc("/getAllType", tasktype.GetAllTasksType(srv))

	mylogger := mylogger.NewMyLogger()
	if errServ := serverStart.NewServer(mylogger); errServ != nil{
		mylogger.Error("ошибка при попытке запуска сервера")
	}
}
