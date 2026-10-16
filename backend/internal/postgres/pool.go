// Package postgres manages PostgreSQL connections and persistence.
package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

// Open creates a shared pool and verifies connectivity. The caller owns the
// pool and must close it after requests have finished. Use a bounded context.
func Open(ctx context.Context, databaseURL string) (*pgxpool.Pool, error) {
	if databaseURL == "" {
		return nil, errors.New("DATABASE_URL is required")
	}
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		// Connection parsing errors may contain credentials; don't expose the URL.
		return nil, errors.New("invalid PostgreSQL connection configuration")
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, errors.New("PostgreSQL connectivity check failed; check DATABASE_URL and database availability")
	}
	return pool, nil
}
