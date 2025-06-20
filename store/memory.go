package store

import (
	"errors"
	"url-shortener/models"
)

var urlDB = make(map[string]models.URL)

func SaveURL(url models.URL) {
	urlDB[url.ID] = url
}

func GetURL(id string) (models.URL, error) {
	url, ok := urlDB[id]
	if !ok {
		return models.URL{}, errors.New("URL not found")
	}
	return url, nil
}
