package rounds

import "testing"

func TestTotalStrokes(t *testing.T) {
	round := Round{
		CourseName: "Visättra",
		Holes: []HoleScore{
			{HoleNumber: 1, Par: 3, Strokes:2},
			{HoleNumber: 2, Par: 3, Strokes:3},
		},
	}

	total := TotalStrokes(round)
	if total != 5 {
		t.Fatalf("got %d strokes", total)
	}
}