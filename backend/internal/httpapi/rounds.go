package httpapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jesrnsu/jesrnsu.github.io/internal/rounds"
)

type roundTotalResponse struct {
	TotalStrokes       int `json:"total_strokes"`
	TotalPar           int `json:"total_par"`
	ScoreRelativeToPar int `json:"score_relative_to_par"`
}

type roundParResponse struct {
	TotalPar int `json:"total_par"`
}

func saveUserRound(service *rounds.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		if service == nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"error": "round storage is unavailable"})
			return
		}
		var round rounds.Round
		if err := c.ShouldBindJSON(&round); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "request body must match the round JSON structure"})
			return
		}
		if err := rounds.ValidateRound(round); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := service.SaveRound(c.Request.Context(), round); err != nil {
			_ = c.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not save round"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "round saved"})
	}
}

func calculateRoundTotal(c *gin.Context) {
	var round rounds.Round
	if err := c.ShouldBindJSON(&round); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "request body must match the round JSON structure"})
		return
	}

	// Validate the completed round.
	if err := rounds.ValidateRound(round); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Translate the result into an HTTP response.
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
