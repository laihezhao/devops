package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "hello world!")
}

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe("127.0.0.1:8080", nil)
	if err != nil {
		panic(err)
	}
}
