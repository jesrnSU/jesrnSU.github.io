// Package rounds will contain round data and scoring rules.
package rounds

type HoleScore struct {
	HoleNumber int `json:"hole_number"`
	Par int `json:"par"`
	Strokes int `json:"strokes"`
}

type Round struct {
	CourseName string `json:"course_name"`
	Holes []HoleScore `json:"holes"`
}


// TODO: Define the round and hole-score types you need for your first feature.
// Decide identifiers, timestamps, player relationships, and score fields yourself.
