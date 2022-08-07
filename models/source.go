package models

import (
	"time"

	"github.com/google/uuid"
)

type Source struct {
	ID            uuid.UUID  `json:"id"`
	Name          string     `json:"name"`
	Author        string     `json:"author"`
	Description   *string    `json:"description"`
	Links         []string   `json:"links"`
	Language      *string    `json:"language"`
	LastUpdated   *time.Time `json:"lastUpdated"`
	ApplicationId uuid.UUID  `json:"applicationId"`
}
