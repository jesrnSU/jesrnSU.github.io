package rounds

import (
	"fmt"
	"strings"
	"time"
)

// TotalStrokes is a pure calculation: it needs no server or database.
// The HTTP handler calls this function, but tests can also call it directly.
// It adds the supplied values without validating them; call ValidateRound before using it with user input.
func TotalStrokes(round Round) int {
	total := 0

	for _, hole := range round.Holes {
		total += hole.Strokes
	}

	return total
}

func TotalPar(round Round) int {
	total := 0

	for _, hole := range round.Holes {
		total += hole.Par
	}

	return total
}

func ScoreRelativeToPar(round Round) int {
	return TotalStrokes(round) - TotalPar(round)
}

// ConvertToStringDate is your unfinished exercise and is not used by the example.
func ConvertToStringDate(round Round) string {
	return "not implemented"
}

// ValidateRound checks a completed round and returns the first validation error.
func ValidateRound(round Round) error {
	if strings.TrimSpace(round.CourseName) == "" {
		return fmt.Errorf("course_name is required")
	}
	if _, err := time.Parse("2006-01-02", round.Date); err != nil {
		return fmt.Errorf("date must be a valid date in YYYY-MM-DD format")
	}
	if len(round.Holes) == 0 {
		return fmt.Errorf("round must contain at least one hole")
	}
	seen := make(map[int]bool)
	for _, hole := range round.Holes {
		if hole.HoleNumber <= 0 {
			return fmt.Errorf("hole_number must be positive")
		}
		if seen[hole.HoleNumber] {
			return fmt.Errorf("hole %d appears more than once", hole.HoleNumber)
		}
		seen[hole.HoleNumber] = true
		if hole.Par <= 0 {
			return fmt.Errorf("hole %d: par must be positive", hole.HoleNumber)
		}
		if hole.Strokes <= 0 {
			return fmt.Errorf("hole %d: strokes must be positive", hole.HoleNumber)
		}
	}
	return nil
}

// TODO: Implement your first save operation here.
// Define a small storage interface here when the operation needs persistence.
// Accept context.Context for operations that call storage.
// Keep this package independent of Gin and PostgreSQL.
