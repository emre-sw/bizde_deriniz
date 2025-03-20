package service

import (
	"auth/internal/domain"
	"auth/internal/service/dto"
	"context"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	GetAllUsers() ([]domain.Auth, error)
	Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error)
	Login(ctx context.Context, req *dto.LoginRequest, c *gin.Context) (*dto.LoginResponse, error)
	ChangePassword(ctx context.Context, req *dto.ChangePasswordRequest) error
	ForgotPassword(ctx context.Context, req *dto.ForgotPasswordRequest) error
	VerifyEmail(ctx context.Context, req *dto.VerifyEmailRequest) (*dto.VerifyEmailResponse, error)
	Logout(ctx context.Context, req *dto.LogoutRequest) error
}
