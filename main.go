package main

import (
	"febri-rss/controllers"
	"febri-rss/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	feeds := r.Group("feeds")
	feeds.GET("", controllers.FindFeeds)
	feeds.POST("", controllers.CreateFeed)
	feeds.DELETE(":id", controllers.DeleteFeed)
	feeds.PATCH("", controllers.PatchFeed)

	r.GET("sent_items", controllers.SentItems)

	r.Run(":2137")
}
