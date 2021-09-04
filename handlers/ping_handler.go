package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func PingHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}
