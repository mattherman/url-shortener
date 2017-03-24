package config

import "os"

// Config contains application configuration information
type Config struct {
	RedisHost string
	Port      string
}

// Read info from config file
func Read() Config {
	redisHost := os.Getenv("REDIS_HOST")
	port := os.Getenv("URL_SH_PORT")

	if redisHost == "" {
		redisHost = "127.0.0.1:6379"
	}

	if port == "" {
		port = "8080"
	}

	return Config{
		RedisHost: redisHost,
		Port:      port,
	}
}
