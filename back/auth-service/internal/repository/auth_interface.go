package repository

import (
	"auth/internal/domain"
	"context"
)

type AuthRepository interface {
	CreateAuth(ctx context.Context, auth *domain.Auth) error
	GetUserByEmail(ctx context.Context, email string) (*domain.Auth, error)
	GetAllUsers() ([]domain.Auth, error)
	UpdatePassword(ctx context.Context, email, password string) error
}