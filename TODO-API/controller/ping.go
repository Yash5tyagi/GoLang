package controller

import (
	"encoding/json"
	"net/http"
	structs "todo-api/views"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := structs.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(data)
		}
	}
}
