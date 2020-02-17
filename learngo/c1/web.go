package main

import (
	"fmt"
	"istio.io/pkg/log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "your request: %v\n", r.URL)
}

func main() {
	http.HandleFunc("/", handler)
	if err := http.ListenAndServe("localhost:8000", nil); err != nil {
		log.Fatal("err:%v\n", err)
	}
}