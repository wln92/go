package main

import (
	"io"
	"log"
	"net/http"
)

type a struct {
}

func (*a) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.String()
	io.WriteString(w, path)
}

func main() {
	err := http.ListenAndServe(":8080", &a{})
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
