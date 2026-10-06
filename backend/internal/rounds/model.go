// Package rounds contains data and calculations independent of HTTP and SQL.
package rounds

// HoleScore describes one hole. JSON tags map Go fields to request field names.
type HoleScore struct {
	HoleNumber int `json:"hole_number"`
	Par        int `json:"par"`
	Strokes    int `json:"strokes"`
}

// Round is your existing model. The example passes Date through as text;
// choosing a date format and validating it remain separate exercises.
type Round struct {
	CourseName string      `json:"course_name"`
	Date       string      `json:"date"`
	Holes      []HoleScore `json:"holes"`
}

// TODO: Decide validation rules before adding persistence.
