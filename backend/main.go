package main

import (
	"log"
	"net/http"
	"time"

  "github.com/gin-gonic/gin"
)

// TelemetryPayload represents the incoming data structure from a smartwatch.
// The `json:"..."` tags tell Go how to map incoming JSON keys to our struct fields.
type TelemetryPayload struct {
	DeviceID  string    `json:"device_id" binding:"required"`
	HeartRate int       `json:"heart_rate" binding:"required"`
	Steps     int       `json:"steps"`
	Timestamp time.Time `json:"timestamp" binding:"required"`
}

func main() {
	// Initialize a default Gin router (comes with logging and recovery middleware built-in)
	r := gin.Default()

	// 1. Health check endpoint (for DevOps monitoring)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	// 2. Data ingestion endpoint for the smartwatch metrics
	r.POST("/api/v1/telemetry", func(c *gin.Context) {
		var payload TelemetryPayload

		// Bind JSON payload to our struct and validate required fields
		// If it fails validation, return a 400 Bad Request (similar to a Java ValidationException)
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// LOG THE DATA: For now, we print it to the terminal. Next step is saving it to Postgres!
		log.Printf("[TELEMETRY RECIEVED] Device: %s | HR: %d | Steps: %d", 
			payload.DeviceID, payload.HeartRate, payload.Steps)

		// Respond to the "smartwatch" with a 201 Created status
		c.JSON(http.StatusCreated, gin.H{
			"message": "Data logged successfully",
			"received_at": time.Now(),
		})
	})

	// Start the server on port 8080
	log.Println("Backend server starting on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
