package services

import (
	"febri-rss/common"
	"febri-rss/controllers"
	"log"
	"time"

	"github.com/onatm/clockwerk"
)

type PurgeNotUpdatedFeedsJob struct{}

func (j PurgeNotUpdatedFeedsJob) Run() {
	log.Default().Println("Running periodic feed purge...")
	common.EnqueueJob(func() {
		err := controllers.PurgeNotUpdatedFeeds(180)
		if err != nil {
			log.Default().Printf("WARNING: Unable to purge not updated feeds: %s", err)
			return
		}
		log.Default().Printf("Finished purging feeds not updated for over %d days", 180)
	})
}

func StartPurgeNotUpdatedFeedsService() {
	var job PurgeNotUpdatedFeedsJob
	c := clockwerk.New()
	c.Every(time.Hour * 24 * 30).Do(job)
	c.Start()

	// Schedule job first time
	common.EnqueueJob(job.Run)
	log.Default().Println("'Purge not updated feeds' service set-up!")
}
