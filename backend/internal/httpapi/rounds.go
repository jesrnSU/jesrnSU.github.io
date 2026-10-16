package httpapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jesrnsu/jesrnsu.github.io/internal/rounds"
)

// roundTotalResponse belongs to the HTTP layer: it defines this endpoint's JSON.
type roundTotalResponse struct {
	TotalStrokes       int `json:"total_strokes"`
	TotalPar           int `json:"total_par"`
	ScoreRelativeToPar int `json:"score_relative_to_par"`
}

type roundParResponse struct {
	TotalPar int `json:"total_par"`
}

// calculateRoundTotal demonstrates HTTP -> Go model -> calculation -> JSON.
// It does not save a round, so success is 200 OK rather than 201 Created.
func calculateRoundTotal(c *gin.Context) {
	// 1. Decode the request body into the model from the rounds package.
	var round rounds.Round
	if err := c.ShouldBindJSON(&round); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request body must match the round JSON structure"})
		return // Stop here; do not calculate a result from a failed decode.
	}

	// 2. Validate the completed round before calculating.
	if err := rounds.ValidateRound(round); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 3. Translate the result into an HTTP response.
	c.JSON(http.StatusOK, roundTotalResponse{
		TotalStrokes:       rounds.TotalStrokes(round),
		TotalPar:           rounds.TotalPar(round),
		ScoreRelativeToPar: rounds.ScoreRelativeToPar(round),
	})
}

func calculateCoursePar(c *gin.Context) {
	var round rounds.Round
	if err := c.ShouldBindJSON(&round); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request body must match the round JSON structure"})
		return
	}

	if err := rounds.ValidateRound(round); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	total := rounds.TotalPar(round)

	c.JSON(http.StatusOK, roundParResponse{TotalPar: total})
}
