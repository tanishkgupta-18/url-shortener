package service

import "url-shortener/store"

func TrackVisit(id, ip, userAgent string) {
	store.LogVisit(id, ip, userAgent)
}

func GetAnalytics(id string) any {
	return store.GetVisits(id)
}
