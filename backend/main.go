package main

import (
	"context"
	"net/http"
	"time"
	"log"

  "github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type TelemetryPayload struct {
	DeviceID  string    `json:"device_id" binding:"required"`
	HeartRate int       `json:"heart_rate" binding:"required"`
	Steps     int       `json:"steps"`
	Timestamp time.Time `json:"timestamp" binding:"required"`
}

func main() {
	r := gin.Default()

	r.POST("/api/v1/telemetry", func(c *gin.Context) {
		var payload TelemetryPayload

		// Parses incoming watch data
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Opens a direct connection to running Docker Postgres instance
		connStr := "postgres://postgres:password@localhost:5432/postgres"
		conn, err := pgx.Connect(context.Background(), connStr)
		if err != nil {
			log.Printf("REAL DATABASE ERROR: %v\n", err) 
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		defer conn.Close(context.Background())

		_, _ = conn.Exec(context.Background(), `
			CREATE TABLE IF NOT EXISTS telemetry (
				device_id TEXT, heart_rate INT, steps INT, timestamp TIMESTAMPTZ
			);
		`)

		_, err = conn.Exec(context.Background(), 
			"INSERT INTO telemetry (device_id, heart_rate, steps, timestamp) VALUES ($1, $2, $3, $4)",
			payload.DeviceID, payload.HeartRate, payload.Steps, payload.Timestamp,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "DB Save Fail"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"message": "Saved successfully!"})
	})

	r.Run(":8080")
}