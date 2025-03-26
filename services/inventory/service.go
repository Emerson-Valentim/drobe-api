package inventory

import (
	"context"
	"fmt"
	"time"

	"github.com/emersonvalentim/drobe-api"
	"github.com/emersonvalentim/drobe-api/internal/uuid"
)

type Repository interface {
	CreateItem(ctx context.Context, item drobe.Item) error
	GetItem(ctx context.Context, id uuid.UUID, ownerID uuid.UUID) (drobe.Item, error)
	ListItems(ctx context.Context, ownerID uuid.UUID) ([]drobe.Item, error)
	DeleteItem(ctx context.Context, id uuid.UUID, ownerID uuid.UUID) error
	UpdateItemLocation(ctx context.Context, id uuid.UUID, ownerID uuid.UUID, location string) error
}

type FileStore interface {
	GetPresignedUploadURL(ctx context.Context, key string, expiry time.Duration) (string, string, error)
	DeleteObject(ctx context.Context, key string) error
}

type Service struct {
	repo      Repository
	filestore FileStore
}

func NewService(repo Repository, filestore FileStore) *Service {
	return &Service{
		repo:      repo,
		filestore: filestore,
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

func (s *Service) UploadImage(ctx context.Context, id uuid.UUID, ownerID uuid.UUID) (string, error) {
	item, err := s.repo.GetItem(ctx, id, ownerID)
	if err != nil {
		return "", err
	}

	location, uploadURL, err := s.filestore.GetPresignedUploadURL(ctx, fmt.Sprintf("%s/%s", ownerID.String(), id.String()), 5*time.Minute)
	if err != nil {
		return "", err
	}

	// If the item already has an image, delete it
	if item.IsImageUploaded() {
		if err := s.filestore.DeleteObject(ctx, item.Location); err != nil {
			return uploadURL, err
		}
	}

	item.UpdateLocation(location)
	if err = s.repo.UpdateItemLocation(ctx, id, ownerID, location); err != nil {
		return uploadURL, err
	}

	return uploadURL, nil
}
