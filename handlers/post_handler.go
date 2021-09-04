package handlers

import (
	"gin-gonic-trial/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PostRequest helps to marshal the data into the Item struct
type PostRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func PostHandler(db *database.Database) gin.HandlerFunc {
	return func(context *gin.Context) {
		// Bind properties of incoming request into the struct
		data := PostRequest{}
		err := context.Bind(data)
		if err != nil {
			return
		}

		// Add to database
		db.Add(database.Item{
			Title:       data.Title,
			Description: data.Description,
		})

		context.JSON(http.StatusOK, gin.H{
			"message": "Posted",
		})
	}
}
