package services

import (
	"febri-rss/common"
	"febri-rss/controllers"
	"log"

	"github.com/onatm/clockwerk"
)

type FetchRssJob struct{}

func (j FetchRssJob) Run() {
	log.Default().Println("Running periodic rss item fetch...")
	common.EnqueueJob(controllers.FetchRssEntries)
}

func StartFetchRssService(configuration common.FebriRssConfiguration) {
	log.Default().Println("Setting-up job scheduler...")
	// Setup scheduler
	var job FetchRssJob
	c := clockwerk.New()
	c.Every(configuration.Services.FetchRss.Period).Do(job)
	c.Start()

	// Schedule job first time
	common.EnqueueJob(controllers.FetchRssEntries)
	log.Default().Println("Fetch RSS Service set-up!")
}
