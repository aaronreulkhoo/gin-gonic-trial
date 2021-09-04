package handlers

import (
	"gin-gonic-trial/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetHandler uses the generic interface which allows us to call the function without passing in the pointer
func GetHandler(db database.Getter) gin.HandlerFunc {
	return func(context *gin.Context) {
		results := db.ReadAll()
		context.JSON(http.StatusOK, results)
	}
}
