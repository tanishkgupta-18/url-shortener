package models

import "time"

type Visit struct {
	ID         string    `json:"id"`
	AccessedAt time.Time `json:"accessed_at"`
	IP         string    `json:"ip"`
	UserAgent  string    `json:"user_agent"`
}
