package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request)  {
	io.WriteString(w, "hello world!3333333333")
}

func main() {
	http.HandleFunc("/", hello)
	// 阿里云服务器上不能使用127.0.0.1
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		panic(err)
	}
}
