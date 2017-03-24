package store

var urls = map[string]string{
	"ggl": "http://www.google.com",
	"so":  "http://www.stackoverflow.com",
	"ds":  "http://www.directsupply.com",
}

// Get returns the URL associated with the alias
func Get(alias string) string {
	return urls[alias]
}

// Set persists an alias with the specified URL
func Set(alias string, url string) {
	urls[alias] = url
}
