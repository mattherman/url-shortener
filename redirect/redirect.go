package redirect

import (
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/mattherman/url-shortener/httputil"
	"github.com/mattherman/url-shortener/store"
)

// Redirect will redirect the client based on the shortened URL provided
func Redirect(w http.ResponseWriter, r *http.Request) {
	alias := mux.Vars(r)["alias"]
	redirectURL := store.Get(alias)
	http.Redirect(w, r, redirectURL, 303)
}

// AddRedirect will create a new alias to the specified URL
func AddRedirect(w http.ResponseWriter, r *http.Request) {
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

	hash := hash(body)

	err = store.Set(hash, bodyString)

	if err != nil {
		httputil.RespondWithError(w, err, http.StatusConflict)
	}

	shortenedURL := "http://" + r.Host + "/" + hash
	httputil.RespondWithValue(w, shortenedURL)
}

func hash(s []byte) string {
	h := fnv.New32a()
	h.Write(s)
	return fmt.Sprint(h.Sum32())
}
