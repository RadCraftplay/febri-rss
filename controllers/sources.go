package controllers

import (
	"bytes"
	"encoding/json"
	"febri-rss/models"
	"fmt"
)

const (
	febri_server_host = "http://localhost"
	febri_server_port = 5286
)

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
