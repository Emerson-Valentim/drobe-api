package clothing

import (
	"context"

	"github.com/emersonvalentim/drobe-api"
	"github.com/emersonvalentim/drobe-api/internal/postgres"
	"github.com/emersonvalentim/drobe-api/internal/uuid"
	"github.com/emersonvalentim/drobe-api/services/clothing/models"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	db *models.Queries
}

func NewRepository(pg *postgres.Postgres) *Repository {
	return &Repository{db: models.New(pg.Conn)}
}

func asCloth(cloth models.Clothe) (drobe.Cloth, error) {
	id, err := uuid.Parse(cloth.ID.String())
	if err != nil {
		return drobe.Cloth{}, err
	}

	return drobe.Cloth{
		ID:        id,
		Name:      cloth.Name,
		Category:  cloth.Category,
		Color:     cloth.Color,
		Size:      cloth.Size,
		Brand:     cloth.Brand,
		Location:  cloth.Location,
		CreatedAt: cloth.CreatedAt.Time,
		UpdatedAt: cloth.UpdatedAt.Time,
	}, nil
}

func (r *Repository) CreateCloth(ctx context.Context, cloth drobe.Cloth) error {
	_, err := r.db.CreateCloth(ctx, models.CreateClothParams{
		ID:        pgtype.UUID{Bytes: cloth.ID.Bytes(), Valid: true},
		Name:      cloth.Name,
		Category:  cloth.Category,
		Color:     cloth.Color,
		Size:      cloth.Size,
		Brand:     cloth.Brand,
		Location:  cloth.Location,
		CreatedAt: pgtype.Timestamp{Time: cloth.CreatedAt.UTC(), Valid: true},
		UpdatedAt: pgtype.Timestamp{Time: cloth.UpdatedAt.UTC(), Valid: true},
	})

	return err
}

func (r *Repository) GetCloth(ctx context.Context, id uuid.UUID) (drobe.Cloth, error) {
	cloth, err := r.db.GetClothByID(ctx, pgtype.UUID{Bytes: id.Bytes(), Valid: true})
	if err != nil {
		return drobe.Cloth{}, err
	}
	return asCloth(cloth)
}

func (r *Repository) ListClothes(ctx context.Context) ([]drobe.Cloth, error) {
	clothes, err := r.db.ListClothes(ctx)
	if err != nil {
		return nil, err
	}

	cloths := make([]drobe.Cloth, 0, len(clothes))
	for _, cloth := range clothes {
		converted, err := asCloth(cloth)
		if err != nil {
			return nil, err
		}
		cloths = append(cloths, converted)
	}

	return cloths, nil
}

func (r *Repository) DeleteCloth(ctx context.Context, id uuid.UUID) error {
	return r.db.DeleteCloth(ctx, pgtype.UUID{Bytes: id.Bytes(), Valid: true})
}
