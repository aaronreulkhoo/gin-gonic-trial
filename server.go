package main

import (
	"fmt"
	"gin-gonic-trial/database"
	"gin-gonic-trial/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Creating db")
	db := database.CreateDatabase()

	fmt.Println("Starting server")
	server := gin.Default()
	server.GET("/ping", handlers.PingHandler())
	server.GET("./", handlers.GetHandler(db))
	server.POST("./", handlers.PostHandler(db))
	err := server.Run(":8080")
	if err != nil {
		fmt.Println("Server failed to start")
		return
	}
}
