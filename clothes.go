package drobe

import "time"

// ClothingItem represents a single piece of clothing
type ClothingItem struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Color     string    `json:"color"`
	Size      string    `json:"size"`
	Brand     string    `json:"brand"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ClothingService defines the interface for clothing-related operations
type ClothingService interface {
	GetItem(id string) (*ClothingItem, error)
	ListItems() ([]ClothingItem, error)
	CreateItem(item *ClothingItem) error
	UpdateItem(item *ClothingItem) error
	DeleteItem(id string) error
}
