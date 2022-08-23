package services

import (
	"febri-rss/common"
	"log"
)

func SetupServices() {
	log.Default().Println("Starting background worker...")
	common.CreateWorker()
	log.Default().Println("Background worker started!")

	log.Default().Println("Starting services...")
	StartFetchRssService()
	log.Default().Println("Services started!")
}
