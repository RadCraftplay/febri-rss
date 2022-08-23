package services

import (
	"febri-rss/common"
	"febri-rss/controllers"
	"log"
	"time"

	"github.com/onatm/clockwerk"
)

type FetchRssJob struct{}

func (j FetchRssJob) Run() {
	log.Default().Println("Running periodic rss item fetch...")
	common.EnqueueJob(controllers.FetchRssEntries)
}

func StartFetchRssService() {
	log.Default().Println("Setting-up job scheduler...")
	// Setup scheduler
	var job FetchRssJob
	c := clockwerk.New()
	c.Every(6 * time.Hour).Do(job)
	c.Start()

	// Schedule job first time
	common.EnqueueJob(controllers.FetchRssEntries)
	log.Default().Println("Fetch RSS Service set-up!")
}
