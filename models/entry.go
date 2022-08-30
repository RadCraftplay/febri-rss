package models

import (
	"time"

	"github.com/google/uuid"
)

type Entry struct {
	ID          uuid.UUID  `json:"id"`
	SourceId    uuid.UUID  `json:"sourceId"`
	Title       string     `json:"title"`
	Links       []string   `json:"links"`
	GUID        string     `json:"guid"`
	Description *string    `json:"description"`
	PubDate     *time.Time `json:"pubDate"`
}
