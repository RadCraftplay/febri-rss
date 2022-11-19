package services

import (
	"febri-rss/common"
	"log"
)

func SetupServices(configuration common.FebriRssConfiguration) {
	log.Default().Println("Starting background worker...")
	common.CreateWorker()
	log.Default().Println("Background worker started!")

	log.Default().Println("Starting services...")
	StartPurgeNotUpdatedFeedsService(configuration)
	StartFetchRssService(configuration)
	log.Default().Println("Services started!")
}
