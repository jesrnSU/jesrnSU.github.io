// Package httpapi translates between HTTP requests and application services.
package httpapi

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/healthz", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:4321")

		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Route registration connects a method and path to a handler function.
	// This learning endpoint calculates a total without touching the database.
	router.POST("/api/v1/rounds/total", calculateRoundTotal)
	router.POST("/api/v1/rounds/course_par", calculateCoursePar)
	// TODO: Accept handler dependencies and register your routes here.
	return router
}
