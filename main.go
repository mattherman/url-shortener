package main

import (
	"log"
	"net/http"

	"github.com/mattherman/url-shortener/redirect"
)

func main() {
	http.HandleFunc("/", redirect.Redirect)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
