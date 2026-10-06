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
