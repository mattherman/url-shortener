package store

import (
	"log"
	"strings"

	"github.com/garyburd/redigo/redis"
)

// NotConnected represents a lack of connection to the data store
type NotConnected struct {
	message string
}

func (e NotConnected) Error() string {
	return e.message
}

// AliasAlreadyExists represents an error where an alias with
// the specified key already exists
type AliasAlreadyExists struct {
	alias string
}

func (e AliasAlreadyExists) Error() string {
	return "An alias with key " + e.alias + " already exists for a different URL."
}

var connection redis.Conn

// CreateConnection will set the Redis connection for the store
func CreateConnection(host string) {
	conn, err := redis.Dial("tcp", host)

	if err != nil {
		log.Fatal("Failed to open connection to Redis: " + err.Error())
		return
	}

	connection = conn
}

// DestroyConnection will close a connection if it exists
func DestroyConnection() {
	err := verifyConnection()

	if err == nil {
		connection.Close()
	}
}

func verifyConnection() error {
	if connection == nil {
		return NotConnected{"No connection to Redis exists. Please connect before accessing the data store. "}
	}

	return nil
}

// Get returns the URL associated with the alias
func Get(alias string) (string, error) {
	err := verifyConnection()

	if err != nil {
		return "", err
	}

	n, err := redis.String(connection.Do("GET", alias))

	if err != nil {
		return "", err
	}

	return n, nil
}

// Set persists an alias with the specified URL. Returns an error
// if the alias already exists.
func Set(alias string, url string) error {
	err := verifyConnection()

	if err != nil {
		return err
	}

	if checkForExistingAlias(alias, url) {
		return AliasAlreadyExists{alias}
	}

	_, err = connection.Do("SET", alias, url)

	if err != nil {
		return err
	}

	return nil
}

func checkForExistingAlias(alias string, url string) bool {
	existingURLMapping, _ := Get(alias)

	return existingURLMapping != "" && strings.ToLower(existingURLMapping) != strings.ToLower(url)
}
