package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 4001

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Starting server on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), mux)
	log.Fatal(err)
}
