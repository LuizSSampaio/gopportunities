package router

import "github.com/gin-gonic/gin"

// Init initializes the router
func Init() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve frontend static files
	router.GET("/api/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start and run the server
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
