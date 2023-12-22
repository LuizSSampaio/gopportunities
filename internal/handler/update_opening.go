package handler

import "github.com/gin-gonic/gin"

func UpdateOpeningHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "PUT opening",
	})
}
