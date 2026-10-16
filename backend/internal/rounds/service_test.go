package rounds

import "testing"

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
