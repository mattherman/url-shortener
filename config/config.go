package config

import "os"

// Config contains application configuration information
type Config struct {
	RedisHost string
}

// Read info from config file
func Read() Config {
	redisHost := os.Getenv("REDIS_HOST")

	if redisHost == "" {
		redisHost = "http://127.0.0.1:6379"
	}

	return Config{RedisHost: redisHost}
}
