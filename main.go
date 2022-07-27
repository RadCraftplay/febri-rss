package main

import (
	"febri-rss/controllers"
	"febri-rss/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.DELETE("/books", controllers.DeleteBook)

	r.Run(":2137")
}
