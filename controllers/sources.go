package controllers

import (
	"bytes"
	"encoding/json"
	"febri-rss/models"
	"fmt"

	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
)

func FetchSourceInfo(url string) (*models.Source, error) {
	parser := gofeed.NewParser()
	data, err := parser.ParseURL(url)
	if err != nil {
		return nil, err
	}

	var author = ""

	if data.Author == nil {
		author = "febri-rss"
	}

	if author == "" {
		author = data.Author.Name
		if author == "" {
			author = data.Author.Email
		}
	}

	return &models.Source{
		ID:            uuid.New(),
		Name:          data.Title,
		Author:        author,
		Description:   &data.Description,
		Links:         []string{url, data.Link},
		Language:      &data.Language,
		ApplicationId: uuid.MustParse("405f4499-6e46-442e-8e14-a59f6733ed26"),
	}, nil
}

type InvalidStatusCodeError struct {
	expectedStatusCode, actualStatusCode int
}

func (e *InvalidStatusCodeError) Error() string {
	return fmt.Sprintf(
		"Invalid status code. Got %d, expected %d",
		e.actualStatusCode,
		e.expectedStatusCode)
}

func CreateSource(s models.Source) error {
	serialized, err := json.Marshal(s)

	if err != nil {
		return err
	}

	buff := bytes.NewBuffer(serialized)

	resp, err := FebriApiClient.Post(
		fmt.Sprintf("%s:%d/api/sources", febri_server_host, febri_server_port),
		"application/json",
		buff)

	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		return &InvalidStatusCodeError{
			expectedStatusCode: 201,
			actualStatusCode:   resp.StatusCode,
		}
	}

	return nil
}
