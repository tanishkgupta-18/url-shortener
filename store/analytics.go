package store

import "time"
import "url-shortener/models"

var analyticsDB = make(map[string][]models.Visit)

func LogVisit(id, ip, userAgent string) {
	visit := models.Visit{
		ID:         id,
		AccessedAt: time.Now(),
		IP:         ip,
		UserAgent:  userAgent,
	}
	analyticsDB[id] = append(analyticsDB[id], visit)
}

func GetVisits(id string) []models.Visit {
	return analyticsDB[id]
}
