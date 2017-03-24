package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mattherman/url-shortener/config"
	"github.com/mattherman/url-shortener/redirect"
)

func main() {
	log.SetOutput(os.Stdout)

	r := mux.NewRouter()

	r.HandleFunc("/{alias}", redirect.Redirect).Methods("GET")
	r.HandleFunc("/create/{alias}", redirect.AddRedirect).Methods("POST")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	logConfig()

	port := "8080"

	log.Println("Ready to accept requests on port " + port + "...")

	log.Fatal(http.ListenAndServe(":"+port, r))
}

func logConfig() config.Config {
	configuration := config.Read()

	log.Println("RedisHost = " + configuration.RedisHost)

	return configuration
}
