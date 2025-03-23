package inventory

import (
	"context"
	"time"

	"github.com/emersonvalentim/drobe-api"
	"github.com/emersonvalentim/drobe-api/internal/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

type CreateInventoryItem struct {
	Name     string    `json:"name"`
	Category string    `json:"category"`
	Size     string    `json:"size"`
	Color    string    `json:"color"`
	Brand    string    `json:"brand"`
	OwnerID  uuid.UUID `json:"ownerId"`
}

func (s *Service) CreateItem(ctx context.Context, input CreateInventoryItem) (drobe.Item, error) {
	id := uuid.New()
	now := time.Now()

	item := drobe.Item{
		ID:        id,
		OwnerID:   input.OwnerID,
		Name:      input.Name,
		Category:  input.Category,
		Color:     input.Color,
		Size:      input.Size,
		Brand:     input.Brand,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := s.repo.CreateItem(ctx, item)

	return item, err
}

func (s *Service) GetItem(ctx context.Context, id uuid.UUID, ownerID uuid.UUID) (drobe.Item, error) {
	item, err := s.repo.GetItem(ctx, id, ownerID)
	if err != nil {
		return drobe.Item{}, err
	}
	return item, nil
}

func (s *Service) ListItems(ctx context.Context, ownerID uuid.UUID) ([]drobe.Item, error) {
	items, err := s.repo.ListItems(ctx, ownerID)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) DeleteItem(ctx context.Context, id uuid.UUID, ownerID uuid.UUID) error {
	return s.repo.DeleteItem(ctx, id, ownerID)
}
