package redirect

import (
	"fmt"
	"hash/fnv"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
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
	body, _ := ioutil.ReadAll(r.Body)

	hash := hash(body)

	store.Set(hash, string(body[:]))

	shortenedURL := "http://" + r.Host + "/" + hash
	fmt.Fprintf(w, shortenedURL)
}

func hash(s []byte) string {
	h := fnv.New32a()
	h.Write(s)
	return fmt.Sprint(h.Sum32())
}
