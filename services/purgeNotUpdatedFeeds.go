package services

import (
	"febri-rss/common"
	"febri-rss/controllers"
	"log"

	"github.com/onatm/clockwerk"
)

type PurgeNotUpdatedFeedsJob struct {
	removeAfterDays uint
}

func (j PurgeNotUpdatedFeedsJob) Run() {
	log.Default().Println("Running periodic feed purge...")
	common.EnqueueJob(func() {
		err := controllers.PurgeNotUpdatedFeeds(j.removeAfterDays)
		if err != nil {
			log.Default().Printf("WARNING: Unable to purge not updated feeds: %s", err)
			return
		}
		log.Default().Printf("Finished purging feeds not updated for over %d days", j.removeAfterDays)
	})
}

func StartPurgeNotUpdatedFeedsService(configuration common.FebriRssConfiguration) {
	var job PurgeNotUpdatedFeedsJob = PurgeNotUpdatedFeedsJob{
		removeAfterDays: configuration.Services.PurgeNotUpdatedFeeds.PurgeAfterDays,
	}

	c := clockwerk.New()
	c.Every(configuration.Services.PurgeNotUpdatedFeeds.Every).Do(job)
	c.Start()

	// Schedule job first time
	common.EnqueueJob(job.Run)
	log.Default().Println("'Purge not updated feeds' service set-up!")
}
