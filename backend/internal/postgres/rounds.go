// Package postgres will implement persistence using the existing pgx dependency.
package postgres

// TODO: Implement the storage interface required by the rounds service.
// Receive a shared connection pool from cmd/api instead of opening connections
// per request. Use request contexts and parameterized SQL.
// Keep schema changes in migrations rather than request handlers.
