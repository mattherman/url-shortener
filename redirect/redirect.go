package redirect

import (
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/mattherman/url-shortener/httputil"
	"github.com/mattherman/url-shortener/store"
)

// Redirect will redirect the client based on the shortened URL provided
func Redirect(w http.ResponseWriter, r *http.Request) {
	alias := mux.Vars(r)["alias"]
	redirectURL, err := store.Get(alias)

	if err != nil {
		httputil.RespondWithError(w, err, 500)
		return
	}

	log.Println("Redirecting " + alias + " to " + redirectURL)

	http.Redirect(w, r, redirectURL, 303)
}

// AddRedirect will create a new alias to the specified URL
func AddRedirect(w http.ResponseWriter, r *http.Request) {
	alias := mux.Vars(r)["alias"]
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		httputil.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	bodyString := string(body[:])

	_, err = url.ParseRequestURI(bodyString)

	if err != nil {
		httputil.RespondWithErrorMessage(w, "Invalid URL provided", http.StatusBadRequest)
		return
	}

	if alias == "" {
		alias = hash(body)
	}

	err = store.Set(alias, bodyString)

	if err != nil {
		httputil.RespondWithError(w, err, http.StatusConflict)
		return
	}

	shortenedURL := "http://" + r.Host + "/" + alias

	log.Println("Created alias to " + bodyString + " with key " + alias)

	httputil.RespondWithValue(w, shortenedURL)
}

func hash(s []byte) string {
	h := fnv.New32a()
	h.Write(s)
	return fmt.Sprint(h.Sum32())
}
