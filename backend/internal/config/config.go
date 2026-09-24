// Package config reads application configuration from the environment.
package config

import "os"

type Config struct {
	HTTPAddr string
}

// Load keeps the prototype's :8080 default. It does not load .env files.
func Load() Config {
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	// TODO: Add and validate configuration as dependencies are introduced.
	return Config{HTTPAddr: addr}
}
