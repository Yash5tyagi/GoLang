package controller

import (
	"encoding/json"
	"net/http"
	"todo-api/model"
	"todo-api/views"
)

func crud() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data := views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data)
			if err := model.CreateTodo(data.Name, data.Todo); err != nil {
				w.Write([]byte("some error"))
				return
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodGet {
			// data, err := model.ReadAll()
			name := r.URL.Query().Get("name")
			data, err := model.ReadByName(name)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodDelete {
			name := r.URL.Path[1:]
			if err := model.DeleteTodo(name); err != nil {
				w.Write([]byte("some error"))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct {
				Status string `json:status`
			}{"Item deleted"})
		}
	}
}
