package main

import (
	"io"
	"net/http"
)

func sayHello(w http.ResponseWriter, r*http.Request) {
	io.WriteString(w, "hello world")
}

func main()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/h", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "/h")
	})
	mux.HandleFunc("/bye", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "/bye")
	})
	mux.HandleFunc("/hello", sayHello)
	http.ListenAndServe(":8080", mux)
}
