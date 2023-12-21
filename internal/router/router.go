package router

import "github.com/gin-gonic/gin"

// Init initializes the router
func Init() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Initialize the routes
	initializeRoutes(router)

	// Start and run the server
	err := router.Run(":8080")
	if err != nil {
		return
	}
}
