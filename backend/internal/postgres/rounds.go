package postgres

import "github.com/jackc/pgx/v5/pgxpool"

// RoundRepository will store and retrieve rounds using the shared pool.
// It does not own the pool; application startup and shutdown manage its lifetime.
type RoundRepository struct {
	pool *pgxpool.Pool
}

func NewRoundRepository(pool *pgxpool.Pool) *RoundRepository {
	return &RoundRepository{pool: pool}
}

// TODO: Add save, list, and get operations after defining the schema.
// Define the required storage interface in the rounds service package.
// Use request contexts, parameterized SQL, and a transaction to save a round
// together with its hole scores. Apply migrations outside request handlers.
