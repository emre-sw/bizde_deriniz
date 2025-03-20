package handler

import (
	"context"

	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	// ADMIN
	GetUsers(ctx context.Context, c *gin.Context)
	
	// USER
	Register(ctx context.Context, c *gin.Context)
	Login(ctx context.Context, c *gin.Context)
	Logout(ctx context.Context, c *gin.Context)
	VerifyEmail(ctx context.Context, c *gin.Context)
	ForgotPassword(ctx context.Context, c *gin.Context)
	ChangePassword(ctx context.Context, c *gin.Context)
	RefreshToken(ctx context.Context, c *gin.Context)
}
