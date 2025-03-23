package inventory

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/emersonvalentim/drobe-api"
	"github.com/emersonvalentim/drobe-api/internal/postgres"
	"github.com/emersonvalentim/drobe-api/internal/uuid"
	"github.com/emersonvalentim/drobe-api/services/inventory/models"
)

type Repository struct {
	db *models.Queries
}

func NewRepository(pg *postgres.Postgres) *Repository {
	return &Repository{db: models.New(pg.Conn)}
}

func asItem(item models.Inventory) (drobe.Item, error) {
	id, err := uuid.Parse(item.ID.String())
	if err != nil {
		return drobe.Item{}, err
	}

	return drobe.Item{
		ID:        id,
		Name:      item.Name,
		Category:  item.Category,
		Color:     item.Color,
		Size:      item.Size,
		Brand:     item.Brand,
		Location:  item.Location,
		CreatedAt: item.CreatedAt.Time,
		UpdatedAt: item.UpdatedAt.Time,
	}, nil
}

func (r *Repository) CreateItem(ctx context.Context, item drobe.Item) error {
	_, err := r.db.CreateItem(ctx, models.CreateItemParams{
		ID:        pgtype.UUID{Bytes: item.ID.Bytes(), Valid: true},
		OwnerID:   pgtype.UUID{Bytes: item.OwnerID.Bytes(), Valid: true},
		Name:      item.Name,
		Category:  item.Category,
		Color:     item.Color,
		Size:      item.Size,
		Brand:     item.Brand,
		Location:  item.Location,
		CreatedAt: pgtype.Timestamp{Time: item.CreatedAt.UTC(), Valid: true},
		UpdatedAt: pgtype.Timestamp{Time: item.UpdatedAt.UTC(), Valid: true},
	})

	return err
}

func (r *Repository) GetItem(ctx context.Context, id uuid.UUID, ownerID uuid.UUID) (drobe.Item, error) {
	item, err := r.db.GetItemByID(ctx, models.GetItemByIDParams{
		ID:      pgtype.UUID{Bytes: id.Bytes(), Valid: true},
		OwnerID: pgtype.UUID{Bytes: ownerID.Bytes(), Valid: true},
	})
	if err != nil {
		return drobe.Item{}, err
	}
	return asItem(item)
}

func (r *Repository) ListItems(ctx context.Context, ownerID uuid.UUID) ([]drobe.Item, error) {
	items, err := r.db.ListItems(ctx, pgtype.UUID{Bytes: ownerID.Bytes(), Valid: true})
	if err != nil {
		return nil, err
	}

	result := make([]drobe.Item, 0, len(items))
	for _, item := range items {
		converted, err := asItem(item)
		if err != nil {
			return nil, err
		}
		result = append(result, converted)
	}

	return result, nil
}

func (r *Repository) DeleteItem(ctx context.Context, id uuid.UUID, ownerID uuid.UUID) error {
	return r.db.DeleteItem(ctx, models.DeleteItemParams{
		ID:      pgtype.UUID{Bytes: id.Bytes(), Valid: true},
		OwnerID: pgtype.UUID{Bytes: ownerID.Bytes(), Valid: true},
	})
}
