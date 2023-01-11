package controllers

import (
	"encoding/json"
	"febri-rss/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type MessageArgs interface{}

type CreateFeedMessage struct {
	URL string `json:"url" binding:"required"`
}

func PostMessage(c *gin.Context) {
	var message models.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	switch strings.ToLower(message.OperationName) {
	case "get feeds":
		feeds := GetFeeds()
		c.JSON(http.StatusOK, gin.H{"data": feeds})

	case "add feed":

		var cfi CreateFeedMessage
		err := json.Unmarshal([]byte(*message.OperationArgsJson), &cfi)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid arguments provided for operation type 'add_feed'",
			})
		}

		feed, e := AddFeed(cfi.URL)
		if e != nil {
			c.JSON(e.ReturnCode, gin.H{
				"message": e.Message,
			})
			break
		}

		c.JSON(http.StatusOK, gin.H{"data": feed})
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid operation name provided",
		})
	}
}
