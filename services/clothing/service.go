package clothing

import (
	"github.com/emersonvalentim/drobe-api"
)

type clothingService struct {
	// TODO: Add storage mechanism (database, file, etc.)
}

func NewClothingService() drobe.ClothingService {
	return &clothingService{}
}

func (s *clothingService) GetItem(id string) (*drobe.ClothingItem, error) {
	// TODO: Implement
	return nil, nil
}

func (s *clothingService) ListItems() ([]drobe.ClothingItem, error) {
	// TODO: Implement
	return nil, nil
}

func (s *clothingService) CreateItem(item *drobe.ClothingItem) error {
	// TODO: Implement
	return nil
}

func (s *clothingService) UpdateItem(item *drobe.ClothingItem) error {
	// TODO: Implement
	return nil
}

func (s *clothingService) DeleteItem(id string) error {
	// TODO: Implement
	return nil
}
