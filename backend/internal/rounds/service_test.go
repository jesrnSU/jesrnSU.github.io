package rounds

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

type roundStoreFunc func(context.Context, Round) error

func (f roundStoreFunc) SaveRound(ctx context.Context, round Round) error {
	return f(ctx, round)
}

func (f roundStoreFunc) GetRound(context.Context, int64) (Round, error) {
	return Round{}, errors.New("unexpected GetRound call")
}

func TestSaveRound(t *testing.T) {
	round := Round{
		CourseName: "Visättra",
		Date:       "2026-10-19",
		Holes:      []HoleScore{{HoleNumber: 1, Par: 3, Strokes: 2}},
	}

	t.Run("passes round and context to storage", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		calls := 0
		service := NewService(roundStoreFunc(func(gotCtx context.Context, got Round) error {
			calls++
			if gotCtx != ctx {
				t.Error("storage received a different context")
			}
			if !reflect.DeepEqual(got, round) {
				t.Errorf("storage received %+v, want %+v", got, round)
			}
			return nil
		}))
		if err := service.SaveRound(ctx, round); err != nil {
			t.Fatal(err)
		}
		if calls != 1 {
			t.Fatalf("storage called %d times, want 1", calls)
		}
	})

	t.Run("rejects invalid round before storage", func(t *testing.T) {
		service := NewService(roundStoreFunc(func(context.Context, Round) error {
			t.Fatal("storage called for invalid round")
			return nil
		}))
		if err := service.SaveRound(context.Background(), Round{}); err == nil {
			t.Fatal("expected validation error")
		}
	})

	t.Run("preserves storage errors", func(t *testing.T) {
		storageErr := errors.New("database unavailable")
		service := NewService(roundStoreFunc(func(context.Context, Round) error {
			return storageErr
		}))
		err := service.SaveRound(context.Background(), round)
		if !errors.Is(err, storageErr) {
			t.Fatalf("SaveRound() error = %v, want wrapped storage error", err)
		}
		if err.Error() != "save round: database unavailable" {
			t.Fatalf("unexpected error message: %v", err)
		}
	})
}

// A domain test runs the calculation without starting HTTP or PostgreSQL.
func TestTotalStrokes(t *testing.T) {
	round := Round{
		CourseName: "Visättra",
		Holes: []HoleScore{
			{HoleNumber: 1, Par: 3, Strokes: 2},
			{HoleNumber: 2, Par: 3, Strokes: 3},
		},
	}
	if got := TotalStrokes(round); got != 5 {
		t.Fatalf("TotalStrokes() = %d, want 5", got)
	}
}

func TestTotalPar(t *testing.T) {
	round := Round{
		CourseName: "Rudan",
		Holes: []HoleScore{
			{HoleNumber: 1, Par: 4, Strokes: 3},
			{HoleNumber: 2, Par: 4, Strokes: 4},
		},
	}
	if got := TotalPar(round); got != 8 {
		t.Fatalf("TotalPar() = %d, want 8", got)
	}
}

func TestRoundEvenPar(t *testing.T) {
	round := Round{
		CourseName: "Visättra",
		Holes: []HoleScore{
			{HoleNumber: 1, Par: 3, Strokes: 2},
			{HoleNumber: 2, Par: 3, Strokes: 4},
		},
	}
	if got := ScoreRelativeToPar(round); got != 0 {
		t.Fatalf("ScoreRelativeToPar() = %d, want 0", got)
	}
}

func TestRoundUnderPar(t *testing.T) {
	round := Round{
		CourseName: "Visättra",
		Holes: []HoleScore{
			{HoleNumber: 1, Par: 3, Strokes: 2},
			{HoleNumber: 2, Par: 3, Strokes: 3},
		},
	}
	if got := ScoreRelativeToPar(round); got != -1 {
		t.Fatalf("ScoreRelativeToPar() = %d, want -1", got)
	}
}

func TestRoundOverPar(t *testing.T) {
	round := Round{
		CourseName: "Visättra",
		Holes: []HoleScore{
			{HoleNumber: 1, Par: 3, Strokes: 4},
			{HoleNumber: 2, Par: 3, Strokes: 3},
		},
	}
	if got := ScoreRelativeToPar(round); got != 1 {
		t.Fatalf("ScoreRelativeToPar() = %d, want 1", got)
	}
}
