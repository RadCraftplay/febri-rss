package controllers

import (
	"net/http"
	"time"

	"febri-rss/common"
	"febri-rss/models"
)

type CreateFeedInput struct {
	URL string `json:"url" binding:"required"`
}

func GetFeeds() []models.Feed {
	var feeds []models.Feed
	models.DB.Find(&feeds)
	return feeds
}

func AddFeed(feedUrl string) (*models.Feed, *models.FebriRssError) {
	feed := models.Feed{
		URL: feedUrl,
	}

	e := models.FebriRssError{
		InnerError: nil,
		Message:    "",
		ReturnCode: http.StatusOK,
	}

	source, err := FetchSourceInfo(feed.URL)
	if err != nil {
		e.InnerError = err
		e.Message = "Unable to fetch feed info"
		e.ReturnCode = http.StatusInternalServerError
		return nil, &e
	}

	err = CreateSource(*source)
	if err != nil {
		e.InnerError = err
		e.Message = "Unable to create source on a remote (Febri) server"
		e.ReturnCode = http.StatusInternalServerError
		return nil, &e
	}

	feed.SourceId = source.ID

	var id uint
	db := models.DB.Raw("INSERT INTO feeds VALUES (DEFAULT, ?, ?) RETURNING ID", feed.URL, feed.SourceId).Scan(&id)
	if db.Error != nil {
		e.InnerError = db.Error
		e.Message = "Unable to add feed due to internal server error"
		e.ReturnCode = http.StatusInternalServerError
		return nil, &e
	}
	feed.ID = id

	common.EnqueueJob(func() {
		FetchRssEntriesSingleFeed(feed)
	})

	return &feed, nil
}

func UpdateFeedPublishedTime(id uint, time *time.Time) error {
	return models.DB.Exec("UPDATE feeds SET last_updated = ? WHERE id = ?", time, id).Error
}

func PurgeNotUpdatedFeeds(afterDays uint) error {
	return models.DB.Exec("delete from feeds where now() - last_updated > ? * '1 day'::interval", afterDays).Error
}
