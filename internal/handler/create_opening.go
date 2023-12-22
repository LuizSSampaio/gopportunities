package handler

import "github.com/gin-gonic/gin"

func CreateOpeningHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "CREATE opening",
	})
}
