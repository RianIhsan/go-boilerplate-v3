package dto

import "time"

type Permission struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	AccessId  uint64    `json:"access_id"`
	Path      string    `json:"path"`
	Method    string    `json:"method"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
