package drobe

import "time"

// Cloth represents a single piece of clothing
type Cloth struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Color     string    `json:"color"`
	Size      string    `json:"size"`
	Brand     string    `json:"brand"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
