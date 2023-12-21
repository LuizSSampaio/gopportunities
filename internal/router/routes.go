package router

import "github.com/gin-gonic/gin"

func initializeRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1/")
	{
		v1.GET("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "GET opening",
			})
		})

		v1.POST("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "POST opening",
			})
		})

		v1.DELETE("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "DELETE opening",
			})
		})

		v1.PUT("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "PUT opening",
			})
		})

		v1.GET("/openings", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "GET all openings",
			})
		})
	}
}
