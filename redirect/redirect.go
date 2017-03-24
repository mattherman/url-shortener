package redirect

import (
	"html"
	"net/http"

	"github.com/mattherman/url-shortener/store"
)

// Redirect will redirect the client based on the shortened URL provided
func Redirect(w http.ResponseWriter, r *http.Request) {
	path := html.EscapeString(r.URL.Path[1:])
	redirectURL := store.Get(path)
	http.Redirect(w, r, redirectURL, 303)
}
