package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../ds"
	"../repository"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := ds.Postrequest{}
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println(data)
			if err := repository.CreateTODO(data.Name, data.Roll_no); err != nil {
				w.Write([]byte("Some error"))
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")
			//fmt.Println("name : ", name)
			if name == "" {
				data, err := repository.ReadAll()
				if err != nil {
					w.Write([]byte(err.Error()))
				}
				json.NewEncoder(w).Encode(data)

			} else {
				data, err := repository.ReadByName(name)
				if err != nil {
					w.Write([]byte(err.Error()))
				}
				json.NewEncoder(w).Encode(data)

			}
		} else if r.Method == http.MethodDelete {
			name := r.URL.Path[13:]
			fmt.Println("name is : ", name)
			if err := repository.DeleteTODO(name); err != nil {
				w.Write([]byte(err.Error()))
			}
			json.NewEncoder(w).Encode(struct {
				Msg string `json:"msg"`
			}{"Item Deleted"})
		}
	}
}
