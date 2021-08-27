package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting server")

	server := gin.Default()
	server.GET("./", get_handler)
	server.POST("./", post_handler)
	server.Run()
}

func get_handler(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func post_handler(context *gin.Context) {
	fmt.Println(&context)
	context.JSON(200, gin.H{
		"message": "Posting",
	})
}

func error_handler(context *gin.Context) {
	fmt.Println(&context)
	context.JSON(400, gin.H{
		"message": "error",
	})
}
