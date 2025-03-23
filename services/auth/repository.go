package auth

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"

	"github.com/emersonvalentim/drobe-api"
	"github.com/emersonvalentim/drobe-api/internal/postgres"
	"github.com/emersonvalentim/drobe-api/internal/uuid"
	"github.com/emersonvalentim/drobe-api/services/auth/models"
)

type Repository struct {
	db *models.Queries
}

func NewRepository(pg *postgres.Postgres) *Repository {
	return &Repository{db: models.New(pg.Conn)}
}

func asUser(user models.User) (drobe.User, error) {
	id, err := uuid.Parse(user.ID.String())
	if err != nil {
		return drobe.User{}, err
	}

	return drobe.User{
		ID:           id,
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt.Time,
		UpdatedAt:    user.UpdatedAt.Time,
	}, nil
}

func (r *Repository) CreateUser(ctx context.Context, user drobe.User) error {
	_, err := r.db.CreateUser(ctx, models.CreateUserParams{
		ID:           pgtype.UUID{Bytes: user.ID.Bytes(), Valid: true},
		Email:        user.Email,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		PasswordHash: user.PasswordHash,
		CreatedAt:    pgtype.Timestamp{Time: user.CreatedAt.UTC(), Valid: true},
		UpdatedAt:    pgtype.Timestamp{Time: user.UpdatedAt.UTC(), Valid: true},
	})
	return err
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (drobe.User, error) {
	user, err := r.db.GetUserByEmail(ctx, email)
	if err != nil {
		return drobe.User{}, err
	}

	return asUser(user)
}

func (r *Repository) GetUserByID(ctx context.Context, id uuid.UUID) (drobe.User, error) {
	user, err := r.db.GetUserByID(ctx, pgtype.UUID{Bytes: id.Bytes(), Valid: true})
	if err != nil {
		return drobe.User{}, err
	}

	return asUser(user)
}
