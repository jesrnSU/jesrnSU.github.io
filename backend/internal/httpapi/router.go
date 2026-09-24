// Package httpapi translates between HTTP requests and application services.
package httpapi

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/healthz", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// TODO: Accept handler dependencies and register your routes here.
	// No application endpoints are registered yet; requests return 404.
	return router
}
