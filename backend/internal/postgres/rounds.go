// Package postgres will implement persistence using the existing pgx dependency.
package postgres

// This file is intentionally not involved in the total-calculation example.
// Calculating a result does not require saving it. When you add a save operation,
// this package will translate your Go model into SQL rows and back.

// TODO: Implement the storage interface required by the rounds service.
// Receive a shared connection pool from cmd/api instead of opening connections
// per request. Use request contexts and parameterized SQL.
// Keep schema changes in migrations rather than request handlers.
