package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Run server....")
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("content-type", "application/json")
			w.Write([]byte("Hello world"))
		}
	})
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Println("Start server error")
	}
}
