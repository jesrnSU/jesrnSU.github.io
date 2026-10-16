// Package config reads application configuration from the environment.
package config

import "os"

type Config struct {
	HTTPAddr    string
	DatabaseURL string
}

// Load keeps the prototype's :8080 default. It does not load .env files.
func Load() Config {
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	// An empty DATABASE_URL keeps calculation-only development database-free.
	return Config{HTTPAddr: addr, DatabaseURL: os.Getenv("DATABASE_URL")}
}
