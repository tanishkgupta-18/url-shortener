package service

import (
	"time"
	"url-shortener/models"
	"url-shortener/store"
	"url-shortener/utils"
)

func CreateShortURL(originalURL string) string {
	short := utils.GenerateShortURL(originalURL)
	url := models.URL{
		ID:           short,
		OriginalURL:  originalURL,
		ShortURL:     short,
		CreationDate: time.Now(),
	}
	store.SaveURL(url)
	return short
}

func GetOriginalURL(id string) (models.URL, error) {
	return store.GetURLByID(id)
}
