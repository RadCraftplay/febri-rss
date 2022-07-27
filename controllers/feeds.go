package controllers

import (
	"fmt"
	"net/http"

	"febri-rss/models"

	"github.com/gin-gonic/gin"
)

type CreateFeedInput struct {
	URL string `json:"url" binding:"required"`
}

type DeleteObjectInput struct {
	Id uint `json:"id" binding:"required"`
}

// GET /feeds
// Get all feeds
func FindFeeds(c *gin.Context) {
	var feeds []models.Feed
	models.DB.Find(&feeds)

	c.JSON(http.StatusOK, gin.H{"data": feeds})
}

// POST /feeds
// Create new feed
func CreateFeed(c *gin.Context) {
	// Validate input
	var input CreateFeedInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create feed
	feed := models.Feed{
		URL: input.URL,
	}

	db := models.DB.Exec(fmt.Sprintf("INSERT INTO %s VALUES (DEFAULT, ?, null)", feed.TableName()), feed.URL)
	if db.Error != nil {
		// TODO: Return 409 Conflict instead?
		c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": feed})
}

// DELETE /feeds
// Delete feed
func DeleteFeed(c *gin.Context) {
	var input DeleteObjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Delete(&models.Feed{}, input.Id)

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// PATCH /feeds
// Update feed
func PatchFeed(c *gin.Context) {
	var input models.Feed
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.DB.Model(&input).Where("id = ?", input.ID).Updates(&input).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}
