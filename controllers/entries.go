package controllers

import (
	"bytes"
	"encoding/json"
	"febri-rss/models"
	"fmt"
	"net/http"
)

func CreateEntry(e models.Entry) error {
	serialized, err := json.Marshal(e)

	if err != nil {
		return err
	}

	buff := bytes.NewBuffer(serialized)

	resp, err := http.Post(
		fmt.Sprintf("%s:%d/api/entries", febri_server_host, febri_server_port),
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
