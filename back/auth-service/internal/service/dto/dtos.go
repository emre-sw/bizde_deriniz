package dto

// RegisterRequest User registration request
// @description User registration request
type RegisterRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"securepassword"`
}

// RegisterResponse User response token information
// @description User response token information
type RegisterResponse struct {
	Message string `json:"message" example:"verification sent to email"`
}

// LoginRequest User login request
// @description User login request
type LoginRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"securepassword"`
}

// LoginResponse User login response
// @description User login response
type LoginResponse struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1N..."`
	RefreshToken string `json:"refresh_token" example:"def456"`
}

// LogoutRequest User logout request
// @description User logout request
type LogoutRequest struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1N..."`
	RefreshToken string `json:"refresh_token" example:"def456"`
}

// ChangePasswordRequest User change password request
// @description User change password request
type ChangePasswordRequest struct {
	Email       string `json:"email" example:"user@example.com"`
	OldPassword string `json:"old_password" example:"oldpassword"`
	NewPassword string `json:"new_password" example:"newpassword"`
}

// ChangePasswordResponse User change password response
// @description User change password response
type ChangePasswordResponse struct {
	Message string `json:"message" example:"Password updated successfully"`
}

// ForgotPasswordRequest User forgot password request
// @description User forgot password request
type ForgotPasswordRequest struct {
	Email string `json:"email" example:"user@example.com"`
}

// ForgotPasswordResponse User forgot password response
// @description User forgot password response
type ForgotPasswordResponse struct {
	Message string `json:"message" example:"Password reset link sent to email"`
}

// ResetPasswordRequest User reset password request
// @description User reset password request
type ResetPasswordRequest struct {
	Token       string `json:"token" example:"1234567890"`
	NewPassword string `json:"new_password" example:"newpassword"`
}

// RefreshTokenRequest User refresh token request
// @description User refresh token request
type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" example:"def456"`
}

// RefreshTokenResponse User refresh token response
// @description User refresh token response
type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1N..."`
	RefreshToken string `json:"refresh_token" example:"def456"`
}

// VerifyEmailRequest User verify email request
// @description User verify email request
type VerifyEmailRequest struct {
	Email            string `json:"email" example:"user@example.com"`
	VerificationCode string `json:"verification_code" example:"123456"`
}

// VerifyEmailResponse User verify email response
// @description User verify email response
type VerifyEmailResponse struct {
	AccessToken  string `json:"access_token" example:"eyJhbGciOiJIUzI1N..."`
	RefreshToken string `json:"refresh_token" example:"def456"`
}
