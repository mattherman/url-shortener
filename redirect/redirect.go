package redirect

import (
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

func AddRedirect(w http.ResponseWriter, r *http.Request) {

}
