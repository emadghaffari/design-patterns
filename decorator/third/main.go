package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	requestKey = 42
)

func main() {
	http.HandleFunc("/", decorate(handler))
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func decorate(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Header.Add("key", "some values...")
		f(w, r)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("key"))
}
