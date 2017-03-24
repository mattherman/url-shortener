package store

import "strings"

var urls = map[string]string{
	"ggl": "http://www.google.com",
	"so":  "http://www.stackoverflow.com",
	"ds":  "http://www.directsupply.com",
}

// Get returns the URL associated with the alias
func Get(alias string) string {
	return urls[alias]
}

// AliasAlreadyExists represents an error where an alias with
// the specified key already exists
type AliasAlreadyExists struct {
	alias string
}

func (e AliasAlreadyExists) Error() string {
	return "An alias with key " + e.alias + " already exists for a different URL."
}

// Set persists an alias with the specified URL. Returns an error
// if the alias already exists.
func Set(alias string, url string) error {
	if checkForExistingAlias(alias, url) {
		return AliasAlreadyExists{alias}
	}

	urls[alias] = url

	return nil
}

func checkForExistingAlias(alias string, url string) bool {
	existingURLMapping := urls[alias]

	return existingURLMapping != "" && strings.ToLower(existingURLMapping) != strings.ToLower(url)
}
