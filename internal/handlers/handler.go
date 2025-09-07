package handlers

import "net/http"

func NewHandlerServer() *http.ServeMux{
	mux := http.NewServeMux()

	return mux
}