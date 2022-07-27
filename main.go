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
	feeds.DELETE("", controllers.DeleteFeed)

	r.Run(":2137")
}
