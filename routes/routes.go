package routes

import (
	"net/http"
)

func CreateMux() *http.ServeMux{
	mux := http.NewServeMux()


	return mux
}