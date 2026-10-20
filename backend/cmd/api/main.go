package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/jesrnsu/jesrnsu.github.io/internal/config"
	"github.com/jesrnsu/jesrnsu.github.io/internal/httpapi"
	"github.com/jesrnsu/jesrnsu.github.io/internal/postgres"
	"github.com/jesrnsu/jesrnsu.github.io/internal/rounds"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.Load()
	var roundService *rounds.Service
	if cfg.DatabaseURL != "" {
		connectCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		pool, err := postgres.Open(connectCtx, cfg.DatabaseURL)
		cancel()
		if err != nil {
			return err
		}
		defer pool.Close()
		log.Print("PostgreSQL connection pool ready")
		repo := postgres.NewRoundRepository(pool)
		roundService = rounds.NewService(repo)
	}

	server := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           httpapi.NewRouter(roundService),
		ReadHeaderTimeout: 5 * time.Second,
	}
	serverErrors := make(chan error, 1)
	go func() { serverErrors <- server.ListenAndServe() }()

	select {
	case err := <-serverErrors:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return fmt.Errorf("serve HTTP: %w", err)
	case <-ctx.Done():
		// Drain active requests before the shared pool is closed.
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()
			return fmt.Errorf("shut down HTTP server: %w", err)
		}
		return nil
	}
}
