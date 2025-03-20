package service

import (
	"auth/internal/domain"
	"auth/internal/repository"
	"auth/internal/service/dto"
	"auth/pkg/configs"
	customErr "auth/pkg/errors"
	"auth/pkg/jwt"
	"auth/pkg/kafka"
	"auth/pkg/redis"
	"auth/pkg/utils"
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthServiceImpl struct {
	authRepository repository.AuthRepository
	kafkaProducer  *kafka.KafkaProducerService
	redisClient    *redis.RedisClient
	tokenService   jwt.TokenService
}

func NewAuthService(
	repo repository.AuthRepository,
	kafkaProducer *kafka.KafkaProducerService,
	redisClient *redis.RedisClient,
	tokenService jwt.TokenService,
) AuthService {
	return &AuthServiceImpl{
		authRepository: repo,
		kafkaProducer:  kafkaProducer,
		redisClient:    redisClient,
		tokenService:   tokenService,
	}
}

func (a *AuthServiceImpl) GetAllUsers() ([]domain.Auth, error) {
	return a.authRepository.GetAllUsers()
}

func (a *AuthServiceImpl) Register(ctx context.Context, req *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	// validate email and password
	if err := utils.ValidateEmailAndPassword(req.Email, req.Password); err != nil {
		return nil, customErr.BadRequest(err.Error())
	}

	// check if user already exists
	existingUser, err := a.authRepository.GetUserByEmail(ctx, req.Email)
	if err == nil && !errors.Is(err, errors.New("user not found")) {
		return nil, customErr.InternalServerError(fmt.Sprintf("failed to check user existence: %v", err))
	}
	if existingUser != nil {
		return nil, customErr.Conflict("user already exists")
	}

	// uuid
	id, err := utils.GenerateID()
	if err != nil {
		return nil, customErr.InternalServerError("failed to generate uuid")
	}

	// hash password
	hashed, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, customErr.InternalServerError("failed to hash password")
	}

	// create
	auth := &domain.Auth{
		ID:       id,
		Email:    req.Email,
		Password: hashed,
	}

	// set auth data to redis
	err = a.redisClient.SetJSON(ctx, "auth:"+req.Email, auth, 10*time.Minute)
	if err != nil {
		return nil, customErr.InternalServerError("failed to set auth data to redis")
	}

	verification := utils.GenerateVerificationToken()
	// generate verification code

	// set verification code to redis
	err = a.redisClient.Set(ctx, "verification:"+req.Email, verification, 10*time.Minute)
	if err != nil {
		return nil, customErr.InternalServerError("failed to set verification code to redis")
	}

	// kafka message
	err = a.kafkaProducer.SendMessage(
		id,
		fmt.Sprintf(`{"event": "user_created", "email": "%s", "verification_code": "%s"}`,
			req.Email, verification,
		),
	)
	if err != nil {
		return nil, customErr.InternalServerError("failed to send message to kafka")
	}

	return &dto.RegisterResponse{
		Message: "verification sent to email",
	}, nil
}

func (s *AuthServiceImpl) Login(ctx context.Context, req *dto.LoginRequest, c *gin.Context) (*dto.LoginResponse, error) {
	config, err := configs.LoadConfig()
	if err != nil {
		return nil, errors.New("failed to load config")
	}

	// check if ip is banned
	ip := c.ClientIP()
	isBanned, err := s.redisClient.IsIPBanned(ctx, ip)
	if err != nil {
		log.Println("failed to check if ip is banned: ", err)
		return nil, customErr.InternalServerError("failed to check if ip is banned")
	}
	if isBanned {
		log.Println("ip banned: ", ip)
		return nil, customErr.Conflict("ip banned")
	}

	// check if user exists
	user, err := s.authRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		s.redisClient.IncrementLoginAttempts(ctx, ip)
		return nil, customErr.InternalServerError("failed to get user by username")
	}
	if user == nil {
		s.redisClient.IncrementLoginAttempts(ctx, ip)
		return nil, customErr.Conflict("user not found")
	}

	// check if password is correct
	err = utils.ComparePassword(req.Password, user.Password)
	if err != nil {
		return nil, customErr.Conflict("password mismatch")
	}

	// reset login attempts
	s.redisClient.ResetLoginAttempts(ctx, ip)

	accessToken, err := s.tokenService.GenerateAccessToken(user.ID)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	refreshToken, err := s.tokenService.GenerateRefreshToken(user.ID)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	// set refresh token to redis
	err = s.redisClient.Set(ctx, "refresh_token:"+user.ID, refreshToken, config.JwtRefreshExpiration)
	if err != nil {
		return nil, errors.New("failed to set refresh token")
	}

	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthServiceImpl) ChangePassword(ctx context.Context, req *dto.ChangePasswordRequest) error {
	user, err := s.authRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return errors.New("user not found")
	}

	err = utils.ComparePassword(req.OldPassword, user.Password)
	if err != nil {
		return errors.New("compare password")
	}

	hassNewPasswd, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		return errors.New("hash error")
	}

	err = s.authRepository.UpdatePassword(ctx, user.Email, hassNewPasswd)
	if err != nil {
		return errors.New("password unupdated")
	}

	return nil
}

func (s *AuthServiceImpl) ForgotPassword(ctx context.Context, req *dto.ForgotPasswordRequest) error {

	// check if user exists
	user, err := s.authRepository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return errors.New("user not found")
	}

	// generate verification token
	verification := utils.GenerateVerificationToken()

	// kafka message
	err = s.kafkaProducer.SendMessage(
		user.ID,
		fmt.Sprintf(`{"event": "user_created", "email": "%s", "verification_code": "%s"}`,
			req.Email, verification,
		),
	)
	if err != nil {
		log.Println("didn't send message to kafka")
	}

	return nil
}

func (s *AuthServiceImpl) VerifyEmail(ctx context.Context, req *dto.VerifyEmailRequest) (*dto.VerifyEmailResponse, error) {

	config, err := configs.LoadConfig()
	if err != nil {
		return nil, errors.New("failed to load config")
	}

	verificationCode, err := s.redisClient.Get(ctx, "verification:"+req.Email)
	if err != nil {
		return nil, errors.New("failed to get verification code")
	}

	if verificationCode != req.VerificationCode {
		return nil, errors.New("invalid verification code")
	}

	auth := &domain.Auth{}
	err = s.redisClient.GetJSON(ctx, "auth:"+req.Email, auth)
	if err != nil {
		return nil, errors.New("failed to get auth data")
	}

	err = s.authRepository.CreateAuth(ctx, auth)
	if err != nil {
		return nil, errors.New("failed to create auth")
	}

	accessToken, err := s.tokenService.GenerateAccessToken(auth.ID)
	if err != nil {
		return nil, errors.New("failed to generate access token")
	}

	refreshToken, err := s.tokenService.GenerateRefreshToken(auth.ID)
	if err != nil {
		return nil, errors.New("failed to generate refresh token")
	}

	// set refresh token to redis
	err = s.redisClient.Set(ctx, "refresh_token:"+auth.ID, refreshToken, config.JwtRefreshExpiration)
	if err != nil {
		return nil, errors.New("failed to set refresh token")
	}

	return &dto.VerifyEmailResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthServiceImpl) Logout(ctx context.Context, req *dto.LogoutRequest) error {
	log.Println("logout request: ", req)

	config, err := configs.LoadConfig()
	if err != nil {
		return errors.New("failed to load config")
	}

	// check if access token is valid
	claim, err := s.tokenService.ValidateToken(req.AccessToken)
	if err != nil {
		return errors.New("invalid access token")
	}
	log.Println("claim: \n", claim)

	accessTokenExpiration := time.Until(time.Unix(claim.ExpiresAt.Unix(), 0))
	if accessTokenExpiration < 0 {
		return errors.New("access token expired")
	}

	// check if refresh token is valid
	refreshClaim, err := s.tokenService.ValidateToken(req.RefreshToken)
	if err != nil {
		return errors.New("invalid refresh token")
	}
	log.Println("refresh claim: \n", refreshClaim)
	// check if refresh token is stored
	storedRefreshToken, err := s.redisClient.Get(ctx, "refresh_token:"+refreshClaim.UserID)
	if err != nil {
		return errors.New("failed to get refresh token")
	}
	log.Println("stored refresh token: \n", storedRefreshToken)
	log.Println("req refresh token: \n", req.RefreshToken)
	if storedRefreshToken != req.RefreshToken {
		return errors.New("invalid refresh token")
	}

	// blacklist access token
	err = s.redisClient.BlacklistToken(ctx, req.AccessToken, config.JwtAccessExpiration)
	if err != nil {
		return errors.New("failed to blacklist access token")
	}

	// blacklist refresh token
	err = s.redisClient.BlacklistToken(ctx, req.RefreshToken, config.JwtRefreshExpiration)
	if err != nil {
		return errors.New("failed to blacklist refresh token")
	}

	// delete refresh token from redis
	err = s.redisClient.Delete(ctx, "refresh_token:"+refreshClaim.UserID)
	if err != nil {
		return errors.New("failed to delete refresh token")
	}

	return nil
}
