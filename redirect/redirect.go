package redirect

import "net/http"

// Redirect will redirect the client based on the shortened URL provided
func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://www.google.com", 303)
}
