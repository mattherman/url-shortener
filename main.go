package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mattherman/url-shortener/redirect"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/{alias}", redirect.Redirect).Methods("GET")
	r.HandleFunc("/create/{alias}", redirect.AddRedirect).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
