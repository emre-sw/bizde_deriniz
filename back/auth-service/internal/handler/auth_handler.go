package handler

import (
	"auth/internal/service"
	"auth/internal/service/dto"
	customErr "auth/pkg/errors"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandlerImpl struct {
	authService service.AuthService
}

func NewAuthHandler(service service.AuthService) AuthHandler {
	return &AuthHandlerImpl{
		authService: service,
	}
}

// GetUsers godoc
// @Summary Get all users
// @Description Returns a list of all registered auths
// @Tags auth
// @Produce json
// @Success 200 {array} dto.RegisterRequest "List of auths"
// @Failure 500 {object} gin.H
// @Router /auth/auths [get]
func (h *AuthHandlerImpl) GetUsers(ctx context.Context, c *gin.Context) {
	users, err := h.authService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}
	c.JSON(http.StatusOK, users)
}

// Register godoc
// @Summary Register a new user
// @Description Creates a new user account with the provided information
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.RegisterRequest true "User registration information"
// @Success 201 {object} dto.RegisterResponse
// @Failure 400 {object} dto.ErrorResponse "Bad Request - Invalid email or password"
// @Failure 409 {object} dto.ErrorResponse "Conflict - User already exists"
// @Failure 500 {object} dto.ErrorResponse "Internal Server Error"
// @Router /auth/register [post]
func (h *AuthHandlerImpl) Register(ctx context.Context, c *gin.Context) {
	var registerRequest dto.RegisterRequest

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.authService.Register(ctx, &registerRequest)
	if err != nil {
		if appErr, ok := err.(*customErr.CustomError); ok {
			c.JSON(appErr.StatusCode(), gin.H{"error": appErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, response)
}

// Login godoc
// @Summary User login
// @Description Authenticates a user and returns access tokens
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.LoginRequest true "User login credentials"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /auth/login [post]
func (h *AuthHandlerImpl) Login(ctx context.Context, c *gin.Context) {
	var loginRequest dto.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.authService.Login(ctx, &loginRequest, c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

// ChangePassword godoc
// @Summary Change user password
// @Description Changes the password for an existing user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.ChangePasswordRequest true "Password change information"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /auth/change-password [post]
func (h *AuthHandlerImpl) ChangePassword(ctx context.Context, c *gin.Context) {
	var changePasswordReq dto.ChangePasswordRequest

	if err := c.ShouldBindJSON(&changePasswordReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.authService.ChangePassword(ctx, &changePasswordReq)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password updated successfully"})
}

// Logout godoc
// @Summary Logout a user
// @Description Logs out a user and invalidates their access token
// @Tags auth
// @Produce json
// @Success 200 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /auth/logout [post]
func (h *AuthHandlerImpl) Logout(ctx context.Context, c *gin.Context) {
	var logoutRequest dto.LogoutRequest

	if err := c.ShouldBindJSON(&logoutRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.authService.Logout(ctx, &logoutRequest)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// RefreshToken godoc
// @Summary Refresh user access token
// @Description Refreshes the access token for a user
// @Tags auth
// @Produce json
// @Param request body dto.RefreshTokenRequest true "Refresh token request"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /auth/refresh-token [post]
func (h *AuthHandlerImpl) RefreshToken(ctx context.Context, c *gin.Context) {

}

// ForgotPassword godoc
// @Summary Forgot user password
// @Description Sends a password reset email to the user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.ForgotPasswordRequest true "Forgot password request"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /auth/forgot-password [post]
func (h *AuthHandlerImpl) ForgotPassword(ctx context.Context, c *gin.Context) {

}

// VerifyEmail godoc
// @Summary Verify user email
// @Description Verifies a user's email address
// @Tags auth
// @Produce json
// @Param request body dto.VerifyEmailRequest true "Verify email request"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /auth/verify-email [post]
func (h *AuthHandlerImpl) VerifyEmail(ctx context.Context, c *gin.Context) {
	var verifyEmailReq dto.VerifyEmailRequest

	if err := c.ShouldBindJSON(&verifyEmailReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.authService.VerifyEmail(ctx, &verifyEmailReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)

}
