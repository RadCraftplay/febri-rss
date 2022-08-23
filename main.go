package main

import (
	"febri-rss/controllers"
	"febri-rss/models"
	"febri-rss/services"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()
	services.SetupServices()

	feeds := r.Group("feeds")
	feeds.GET("", controllers.FindFeeds)
	feeds.POST("", controllers.CreateFeed)
	feeds.DELETE(":id", controllers.DeleteFeed)
	feeds.PATCH("", controllers.PatchFeed)

	r.Run(":2137")
}
