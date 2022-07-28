package controllers

import (
	"net/http"

	"febri-rss/models"

	"github.com/KonishchevDmitry/go-rss"
	"github.com/gin-gonic/gin"
)

func isElementExist(s []models.SentItems, str string) bool {
	for _, v := range s {
		if v.GUID == str {
			return true
		}
	}
	return false
}

// GET /sent_items
// Get all sent items
func SentItems(c *gin.Context) {
	var items map[string]([]models.SentItems) = make(map[string]([]models.SentItems))
	var feeds []models.Feed
	models.DB.Find(&feeds)

	for _, feed := range feeds {
		/* So the algorithm goes like this:
		 * 1. Get the feed data
		 * 2. Get all sent item info FOR THAT FEED from the database
		 * 3. If guid is not present, add item to the send queue (probably list to send all entries at once)
		 * 4. Send all endies to the api with the POST response
		 * 5. If no errors, add all sources to our database
		 */

		data, err := rss.Get(feed.URL)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		var guids []models.SentItems
		dbc := models.DB.Where("feed_id = ?", feed.ID).Select("guid").Find(&guids)
		if dbc.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": dbc.Error})
			return
		}

		items[feed.URL] = make([]models.SentItems, 0)

		for _, item := range data.Items {
			if isElementExist(guids, item.Guid.Id) {
				continue
			}

			si := models.SentItems{
				Feed_ID: feed.ID,
				GUID:    item.Guid.Id,
			}
			items[feed.URL] = append(items[feed.URL], si)
			models.DB.Create(&si)
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": items})
}
