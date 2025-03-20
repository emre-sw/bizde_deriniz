package jwt

import (
	"auth/pkg/configs"
	"auth/pkg/redis"
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type tokenPayload struct {
	UserID    string
	ExpiresAt time.Time
	Issuer    string
	Subject   string
	Secret    []byte
}

type TokenService interface {
	GenerateAccessToken(userID string) (string, error)
	GenerateRefreshToken(userID string) (string, error)
	ValidateToken(token string) (*claim, error)
}

type tokenService struct {
	redisClient *redis.RedisClient
}

func NewTokenService(redisClient *redis.RedisClient) TokenService {
	return &tokenService{
		redisClient: redisClient,
	}
}

func (s *tokenService) GenerateAccessToken(userID string) (string, error) {
	config, err := configs.LoadConfig()
	if err != nil {
		return "", err
	}

	return generateToken(tokenPayload{
		UserID:    userID,
		ExpiresAt: time.Now().Add(time.Duration(config.JwtAccessExpiration)),
		Issuer:    "auth-service",
		Subject:   "auth-service",
		Secret:    []byte(config.JwtAccessSecret),
	})
}

func (s *tokenService) GenerateRefreshToken(userID string) (string, error) {
	config, err := configs.LoadConfig()
	if err != nil {
		return "", err
	}

	return generateToken(tokenPayload{
		UserID:    userID,
		ExpiresAt: time.Now().Add(time.Duration(config.JwtRefreshExpiration)),
		Issuer:    "auth-service",
		Subject:   "auth-service",
		Secret:    []byte(config.JwtRefreshSecret),
	})
}

func (s *tokenService) ValidateToken(token string) (*claim, error) {
	config, err := configs.LoadConfig()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	isBlacklisted, err := s.redisClient.IsTokenBlacklisted(ctx, token)
	if err != nil {
		return nil, err
	}
	if isBlacklisted {
		return nil, fmt.Errorf("token is blacklisted")
	}

	parsedToken, err := jwt.ParseWithClaims(token, &claim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.JwtAccessSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(*claim); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
