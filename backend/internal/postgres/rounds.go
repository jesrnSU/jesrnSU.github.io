package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jesrnsu/jesrnsu.github.io/internal/rounds"
)

// RoundRepository will store and retrieve rounds using the shared pool.
// It does not own the pool; application startup and shutdown manage its lifetime.
type RoundRepository struct {
	pool *pgxpool.Pool
}

func NewRoundRepository(pool *pgxpool.Pool) *RoundRepository {
	return &RoundRepository{pool: pool}
}

func (r *RoundRepository) SaveRound(
	ctx context.Context,
	round rounds.Round,
) error {
	date, err := time.Parse("2006-01-02", round.Date)
	if err != nil {
		return fmt.Errorf("parse round date: %w", err)
	}

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	roundID, err := r.insertRound(ctx, tx, round.CourseName, date)
	if err != nil {
		return err
	}

	if err := r.insertHoleScores(ctx, tx, roundID, round.Holes); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit round: %w", err)
	}

	return nil
}

func (r *RoundRepository) insertRound(
	ctx context.Context,
	tx pgx.Tx,
	courseName string,
	date time.Time,
) (int64, error) {
	var roundID int64

	err := tx.QueryRow(ctx, `
		INSERT INTO rounds (course_name, date)
		VALUES ($1, $2)
		RETURNING id
	`, courseName, date).Scan(&roundID)
	if err != nil {
		return 0, fmt.Errorf("insert round: %w", err)
	}

	return roundID, nil
}

func (r *RoundRepository) insertHoleScores(
	ctx context.Context,
	tx pgx.Tx,
	roundID int64,
	holes []rounds.HoleScore,
) error {
	for _, hole := range holes {
		_, err := tx.Exec(ctx, `
			INSERT INTO hole_scores
				(round_id, hole_number, par, strokes)
			VALUES ($1, $2, $3, $4)
		`, roundID, hole.HoleNumber, hole.Par, hole.Strokes)
		if err != nil {
			return fmt.Errorf("insert hole %d: %w", hole.HoleNumber, err)
		}
	}

	return nil
}

// TODO: Add list and get operations.
