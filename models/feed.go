package models

import "github.com/google/uuid"

type Feed struct {
	ID       uint      `json:"id" gorm:"primary_key"`
	URL      string    `json:"url"`
	SourceId uuid.UUID `json:"source_id"`
}
