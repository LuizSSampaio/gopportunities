package handler

import "github.com/gin-gonic/gin"

func ShowOpeningHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "GET opening",
	})
}

func CreateOpeningHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "CREATE opening",
	})
}

func DeleteOpeningHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "DELETE opening",
	})
}

func UpdateOpeningHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "PUT opening",
	})
}

func ShowAllOpeningHandler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "GET all opening",
	})
}
