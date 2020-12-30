package controller

import (
	"fmt"
	"net/http"
	"../ds"
	"encoding/json"
)
func SendData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//w.Write([]byte("Hello World!"))
		fmt.Println("request received")
		if(r.Method == http.MethodGet) {
			data := ds.Response{
				Code : http.StatusOK,
				Body : "its working!"   , 
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		}

	}
}