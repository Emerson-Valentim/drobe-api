package drobe

import (
	"time"

	"github.com/emersonvalentim/drobe-api/internal/uuid"
)

// Item represents a single piece of inventory
type Item struct {
	ID        uuid.UUID `json:"id"`
	OwnerID   uuid.UUID `json:"ownerId"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Color     string    `json:"color"`
	Size      string    `json:"size"`
	Brand     string    `json:"brand"`
	Location  string    `json:"location"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
