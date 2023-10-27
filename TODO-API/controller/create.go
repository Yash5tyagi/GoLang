package controller

import (
	"net/http"
	"todo-api/model"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			if err := model.CreateTodo(); err != nil {
				w.Write([]byte("some error"))
				return
			}
		}
	}
}
