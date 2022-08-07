package common

import (
	"febri-rss/models"

	"github.com/KonishchevDmitry/go-rss"
	"github.com/google/uuid"
)

func FetchSourceInfo(url string) (*models.Source, error) {
	data, err := rss.Get(url)
	if err != nil {
		return nil, err
	}

	return &models.Source{
		ID:            uuid.New(),
		Name:          data.Title,
		Author:        "febri-rss",
		Description:   &data.Description,
		Links:         []string{url, data.Link},
		Language:      &data.Language,
		ApplicationId: uuid.MustParse("405f4499-6e46-442e-8e14-a59f6733ed26"),
	}, nil
}
