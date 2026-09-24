package main

import (
	"log"

	"github.com/jesrnsu/jesrnsu.github.io/internal/config"
	"github.com/jesrnsu/jesrnsu.github.io/internal/httpapi"
)

func main() {
	cfg := config.Load()
	// TODO: Initialize shared dependencies here and close them on shutdown.
	// Wire storage into services, then services into HTTP handlers.
	router := httpapi.NewRouter()

	// TODO: Configure an http.Server with timeouts and graceful shutdown
	// when preparing this application for deployment.
	if err := router.Run(cfg.HTTPAddr); err != nil {
		log.Fatal(err)
	}
}
