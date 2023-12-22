package handler

import "github.com/gin-gonic/gin"

func ShowOpeningHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "GET opening",
	})
}
