package models

type Feed struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	URL      string `json:"url"`
	SourceId *uint  `json:"source_id"`
}

func (Feed) TableName() string {
	return "rss_service.feeds"
}
