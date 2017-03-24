package redirect

import (
	"html"
	"net/http"
)

var urls = map[string]string{
	"ggl": "http://www.google.com",
	"so":  "http://www.stackoverflow.com",
	"ds":  "http://www.directsupply.com",
}

// Redirect will redirect the client based on the shortened URL provided
func Redirect(w http.ResponseWriter, r *http.Request) {
	path := html.EscapeString(r.URL.Path[1:])
	http.Redirect(w, r, urls[path], 303)
}
