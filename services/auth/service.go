package auth

import (
	"context"
	"encoding/base64"
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/argon2"

	"github.com/emersonvalentim/drobe-api"
	"github.com/emersonvalentim/drobe-api/internal/uuid"
)

const (
	sessionDuration = time.Minute * 60
)

type Repository interface {
	CreateUser(ctx context.Context, user drobe.User) error
	GetUserByEmail(ctx context.Context, email string) (drobe.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (drobe.User, error)
}

type Service struct {
	repo         Repository
	passwordSalt string
	sessionSalt  string
}

func NewService(passwordSalt, sessionSalt string, repo Repository) *Service {
	return &Service{
		passwordSalt: passwordSalt,
		sessionSalt:  sessionSalt,
		repo:         repo,
	}
}

type SignUp struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

func (s *Service) SignUp(ctx context.Context, input SignUp) (drobe.User, error) {
	now := time.Now()

	user := drobe.User{
		ID:           uuid.New(),
		Email:        input.Email,
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		PasswordHash: s.hashPassword(input.Password),
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return drobe.User{}, err
	}

	return user, nil
}

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Service) SignIn(ctx context.Context, input SignIn) (drobe.Session, error) {
	user, err := s.repo.GetUserByEmail(ctx, input.Email)
	if err != nil {
		return drobe.Session{}, err
	}

	incomingPassword := s.hashPassword(input.Password)

	if user.PasswordHash != incomingPassword {
		return drobe.Session{}, errors.New("failed to sign in")
	}

	session, err := s.createSession(user)
	if err != nil {
		return drobe.Session{}, err
	}

	return session, nil
}

func (s *Service) ValidateToken(ctx context.Context, token string) (uuid.UUID, error) {
	claims, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.sessionSalt), nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}

	id, ok := claims.Claims.(jwt.MapClaims)["sub"].(string)
	if !ok {
		return uuid.UUID{}, errors.New("forbidden")
	}

	userID, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, err
	}

	return userID, nil
}

func (s *Service) hashPassword(password string) string {
	hash := argon2.IDKey([]byte(password), []byte(s.passwordSalt), 1, 64*1024, 4, 32)
	identifier := "argon2id$v=" + strconv.Itoa(argon2.Version)
	return identifier + "$" + base64.RawStdEncoding.EncodeToString(hash)
}

func (s *Service) createSession(user drobe.User) (drobe.Session, error) {
	expiresAt := time.Now().Add(sessionDuration)

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": expiresAt.Unix(),
	}).SignedString([]byte(s.sessionSalt))
	if err != nil {
		return drobe.Session{}, err
	}

	return drobe.Session{Token: token, User: user, ExpiresAt: expiresAt}, nil
}
