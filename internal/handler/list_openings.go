package handler

import "github.com/gin-gonic/gin"

func ListOpeningsHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "GET all opening",
	})
}
