package rounds

func TotalStrokes(round Round) int {
	total := 0

	for _, hole := range round.Holes {
		total += hole.Strokes
	}

	return total;
}

func ConvertToStringDate(round Round) string {
	return "not implemented"
}

// TODO: Implement your first round operation and its validation here.
// Define a small storage interface here when the operation needs persistence.
// Accept context.Context for operations that call storage.
// Keep this package independent of Gin and PostgreSQL.
