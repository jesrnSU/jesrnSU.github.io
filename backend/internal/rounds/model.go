// Package rounds contains data and calculations independent of HTTP and SQL.
package rounds

// HoleScore describes one hole. JSON tags map Go fields to request field names.
type HoleScore struct {
	HoleNumber int `json:"hole_number"`
	Par        int `json:"par"`
	Strokes    int `json:"strokes"`
}

// Round describes a completed round. Date uses YYYY-MM-DD.
type Round struct {
	Id 				 int64 			 `json:"id"`
	CourseName string      `json:"course_name"`
	Date       string      `json:"date"`
	Holes      []HoleScore `json:"holes"`
}

// ValidateRound checks these fields before HTTP calculations.
