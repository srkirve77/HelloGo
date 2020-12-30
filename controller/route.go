package controller

import (
	"net/http"
)

func Ping() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/go", SendData())
	mux.HandleFunc("/create", create())
	mux.HandleFunc("/create/", create())
	return mux
}
