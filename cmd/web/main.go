package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 4001

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Starting server on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	log.Fatal(err)
}
