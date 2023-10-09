package main

import (
	"fmt"
	"io"
	"net/http"
)

type logWriter struct{}

func (logWriter) Write(b []byte) (int, error) {
	cont := string(b)
	fmt.Println(cont)
	fmt.Println("Number of bytes read=", len(cont))
	return len(cont), nil
}

func main() {
	res, err := http.Get("https://google.com")
	if err != nil {
		panic(err)
	}
	lw := logWriter{}
	io.Copy(lw, res.Body)
}
