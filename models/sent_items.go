package models

type SentItems struct {
	ID      *uint  `json:"id" gorm:"primary_key"`
	Feed_ID uint   `json:"feed_id"`
	GUID    string `json:"guid"`
}
