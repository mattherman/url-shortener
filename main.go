package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mattherman/url-shortener/config"
	"github.com/mattherman/url-shortener/redirect"
	"github.com/mattherman/url-shortener/store"
)

func main() {
	log.SetOutput(os.Stdout)

	config := getConfig()

	store.CreateConnection(config.RedisHost)
	defer store.DestroyConnection()

	r := mux.NewRouter()

	r.HandleFunc("/{alias}", redirect.Redirect).Methods("GET")
	r.HandleFunc("/create/{alias}", redirect.AddRedirect).Methods("POST")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	log.Println("Ready to accept requests on port " + config.Port + "...")

	log.Fatal(http.ListenAndServe(":"+config.Port, r))
}

func getConfig() config.Config {
	configuration := config.Read()

	log.Println("RedisHost = " + configuration.RedisHost)
	log.Println("Port = " + configuration.Port)

	return configuration
}
