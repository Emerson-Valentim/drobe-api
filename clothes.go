package drobe

import (
	"time"

	"github.com/emersonvalentim/drobe-api/internal/uuid"
)

// Cloth represents a single piece of clothing
type Cloth struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Color     string    `json:"color"`
	Size      string    `json:"size"`
	Brand     string    `json:"brand"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
