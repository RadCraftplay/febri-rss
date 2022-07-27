package models

import "time"

type Feed struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	URL        string    `json:"url"`
	LastActive time.Time `json:"last_active"`
}

func (Feed) TableName() string {
	return "rss_service.feeds"
}
