package handler

import "github.com/gin-gonic/gin"

func DeleteOpeningHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "DELETE opening",
	})
}
