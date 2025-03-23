package clothing

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

type CreateCloth struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Size     string `json:"size"`
	Color    string `json:"color"`
	Brand    string `json:"brand"`
}

func (s *Service) CreateCloth(ctx context.Context, input CreateCloth) (drobe.Cloth, error) {
	id := uuid.New()
	now := time.Now()

	cloth := drobe.Cloth{
		ID:        id,
		Name:      input.Name,
		Category:  input.Category,
		Color:     input.Color,
		Size:      input.Size,
		Brand:     input.Brand,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := s.repo.CreateCloth(ctx, cloth)

	return cloth, err
}

func (s *Service) GetCloth(ctx context.Context, id uuid.UUID) (drobe.Cloth, error) {
	cloth, err := s.repo.GetCloth(ctx, id)
	if err != nil {
		return drobe.Cloth{}, err
	}
	return cloth, nil
}

func (s *Service) ListClothes(ctx context.Context) ([]drobe.Cloth, error) {
	clothes, err := s.repo.ListClothes(ctx)
	if err != nil {
		return nil, err
	}
	return clothes, nil
}

func (s *Service) DeleteCloth(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteCloth(ctx, id)
}
