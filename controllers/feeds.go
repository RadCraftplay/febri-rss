package controllers

import (
	"net/http"
	"time"

	"febri-rss/common"
	"febri-rss/models"

	"github.com/gin-gonic/gin"
)

type CreateFeedInput struct {
	URL string `json:"url" binding:"required"`
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

	source, err := FetchSourceInfo(input.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	err = CreateSource(*source)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	feed.SourceId = source.ID

	var id uint
	db := models.DB.Raw("INSERT INTO feeds VALUES (DEFAULT, ?, ?) RETURNING ID", feed.URL, feed.SourceId).Scan(&id)
	if db.Error != nil {
		// TODO: Return 409 Conflict instead?
		c.JSON(http.StatusInternalServerError, gin.H{"error": db.Error.Error()})
		return
	}
	feed.ID = id

	common.EnqueueJob(func() {
		FetchRssEntriesSingleFeed(feed)
	})

	c.JSON(http.StatusOK, gin.H{"data": feed})
}

func UpdateFeedPublishedTime(id uint, time *time.Time) error {
	return models.DB.Exec("UPDATE feeds SET last_updated = ? WHERE id = ?", time, id).Error
}

func PurgeNotUpdatedFeeds(afterDays uint) error {
	return models.DB.Exec("delete from feeds where now() - last_updated > ? * '1 day'::interval", afterDays).Error
}
