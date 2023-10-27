package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request recieved")
		w.Write([]byte("hello world\n"))
	})
	http.ListenAndServe("localhost:3000", mux)
}
