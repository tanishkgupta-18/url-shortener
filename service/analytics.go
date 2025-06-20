package service

import (
	"time"
	"url-shortener/models"
	"url-shortener/store"
)

// TrackVisit saves visit info to MongoDB
func TrackVisit(id, ip, userAgent string) error {
	visit := models.Visit{
		ID:        id,
		IP:        ip,
		UserAgent: userAgent,
		AccessedAt: time.Now(),
	}
	return store.SaveVisit(visit) 
}

// GetAnalytics returns all visit records for given URL id
func GetAnalytics(id string) ([]models.Visit, error) {
	return store.GetVisitsByID(id) 
}
