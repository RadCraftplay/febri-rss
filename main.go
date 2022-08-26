package main

import (
	"febri-rss/common"
	"febri-rss/controllers"
	"febri-rss/models"
	"febri-rss/services"

	"github.com/gin-gonic/gin"
)

func main() {
	configuration := common.LoadConfiguration()

	r := gin.Default()

	models.ConnectDatabase()
	controllers.SetupHttpClients(configuration)
	services.SetupServices(configuration)

	feeds := r.Group("feeds")
	feeds.GET("", controllers.FindFeeds)
	feeds.POST("", controllers.CreateFeed)
	feeds.DELETE(":id", controllers.DeleteFeed)
	feeds.PATCH("", controllers.PatchFeed)

	r.Run(":2137")
}
