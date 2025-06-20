package handler

import (
	"encoding/json"
	"net/http"
	"strings"
	"url-shortener/service"
)

func Root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}

func Shorten(w http.ResponseWriter, r *http.Request) {
	var data struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	short := service.CreateShortURL(data.URL)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"short_url": short,
	})
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/redirect/")
	url, err := service.GetOriginalURL(id)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	// Analytics tracking
	ip := r.RemoteAddr
	ua := r.UserAgent()
	service.TrackVisit(id, ip, ua)

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func Analytics(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/analytics/")
	visits, err := service.GetAnalytics(id)
	if err != nil {
		http.Error(w, "Analytics not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(visits)
}
