package rounds

// TotalStrokes is a pure calculation: it needs no server or database.
// The HTTP handler calls this function, but tests can also call it directly.
// It adds the supplied values without validating them; validation is a future step.
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

// ConvertToStringDate is your unfinished exercise and is not used by the example.
func ConvertToStringDate(round Round) string {
	return "not implemented"
}

// TODO: Implement your first round operation and its validation here.
// Define a small storage interface here when the operation needs persistence.
// Accept context.Context for operations that call storage.
// Keep this package independent of Gin and PostgreSQL.
