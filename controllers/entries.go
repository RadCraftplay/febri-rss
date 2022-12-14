package controllers

import (
	"bytes"
	"encoding/json"
	"febri-rss/models"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
)

func CreateEntry(e models.Entry) error {
	serialized, err := json.Marshal(e)

	if err != nil {
		return err
	}

	buff := bytes.NewBuffer(serialized)
	resp, err := FebriApiClient.Post(
		fmt.Sprintf("%s:%d/api/entries", febri_server_host, febri_server_port),
		"application/json",
		buff)

	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		return &InvalidStatusCodeError{
			expectedStatusCode: 201,
			actualStatusCode:   resp.StatusCode,
		}
	}

	return nil
}

func isElementExist(s []models.SentItems, str string) bool {
	for _, v := range s {
		if v.GUID == str {
			return true
		}
	}
	return false
}

func FetchRssEntries() {
	var feeds []models.Feed
	models.DB.Find(&feeds)

	log.Default().Println("Fetching rss entries...")

	for _, feed := range feeds {
		updated_time, time := FetchRssEntriesSingleFeed(feed)
		if updated_time {
			err := UpdateFeedPublishedTime(feed.ID, time)
			if err != nil {
				log.Default().Printf("WARNING: Unable to update feed LastUpdated time: %s", err)
			}
		}
	}

	log.Default().Println("Finished fetching rss entries!")
}

func FetchRssEntriesSingleFeed(feed models.Feed) (bool, *time.Time) {
	/* So the algorithm goes like this:
	 * 1. Get the feed data
	 * 2. Get all sent item info FOR THAT FEED from the database
	 * 3. If guid is not present, add item to the send queue (probably list to send all entries at once)
	 * 4. Send all endies to the api with the POST response
	 * 5. If no errors, add all sources to our database
	 */

	parser := gofeed.NewParser()
	data, err := parser.ParseURL(feed.URL)
	if err != nil {
		log.Default().Printf("WARNING: %s\n", err)
		return false, nil
	}

	var guids []models.SentItems
	dbc := models.DB.Where("feed_id = ?", feed.ID).Select("guid").Find(&guids)
	if dbc.Error != nil {
		log.Default().Printf("WARNING: %s\n", dbc.Error)
		return false, nil
	}

	var lastUpdated *time.Time = nil

	for _, item := range data.Items {
		if isElementExist(guids, item.GUID) {
			continue
		}

		e := models.Entry{
			ID:          uuid.New(),
			SourceId:    feed.SourceId,
			Title:       item.Title,
			Links:       []string{item.Link},
			GUID:        item.GUID,
			Description: &item.Description,
			PubDate:     item.PublishedParsed,
		}

		err = CreateEntry(e)
		if err != nil {
			log.Default().Printf("WARNING: %s\n", err)
			return lastUpdated != nil, lastUpdated
		}

		if lastUpdated == nil || lastUpdated.Before(*e.PubDate) {
			lastUpdated = e.PubDate
		}

		si := models.SentItems{
			Feed_ID: feed.ID,
			GUID:    item.GUID,
		}
		models.DB.Create(&si)
	}

	return lastUpdated != nil, lastUpdated
}
