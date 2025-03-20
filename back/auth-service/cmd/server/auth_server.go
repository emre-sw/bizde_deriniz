package server

import (
	_ "auth/docs"
	"auth/internal/handler"
	"auth/internal/repository"
	"auth/internal/service"
	"auth/pkg/configs"
	"auth/pkg/db"
	"auth/pkg/jwt"
	"auth/pkg/kafka"
	"auth/pkg/redis"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AuthServer(r *gin.Engine, config *configs.Config) {
	// connect to db
	conn := db.ConnectDB(config)

	// connect to redis
	redisClient, err := redis.NewRedisClient()
	if err != nil {
		log.Println("redis couldn't run")
	}

	// connect to kafka
	kafka, err := kafka.NewKafkaProducerService(config)
	if err != nil {
		log.Println("kafka couldn't run")
	}

	// connect to jwt
	tokenService := jwt.NewTokenService(redisClient)

	// connect to middleware
	middleware := NewAuthMiddleware(tokenService)

	// connect to repository
	authRepo := repository.NewAuthRepository(conn)

	// connect to service
	authService := service.NewAuthService(authRepo, kafka, redisClient, tokenService)

	// connect to handler
	authHandler := handler.NewAuthHandler(authService)

	// Swagger configuration
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler,
		ginSwagger.URL("/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1)))

	// ADMIN
	adminRoutes := r.Group("/admin")
	adminRoutes.Use(middleware.JwtMiddleware())
	{
		adminRoutes.GET("/auths", func(c *gin.Context) { authHandler.GetUsers(c.Request.Context(), c) })
	}

	// auth routes
	authRoutes := r.Group("/auth")
	{
		// USER
		authRoutes.POST("/register", func(c *gin.Context) { authHandler.Register(c.Request.Context(), c) })
		authRoutes.POST("/login", func(c *gin.Context) { authHandler.Login(c.Request.Context(), c) })
		authRoutes.POST("/logout", func(c *gin.Context) { authHandler.Logout(c.Request.Context(), c) })
		authRoutes.POST("/change-password", func(c *gin.Context) { authHandler.ChangePassword(c.Request.Context(), c) })
		authRoutes.POST("/forgot-password", func(c *gin.Context) { authHandler.ForgotPassword(c.Request.Context(), c) })
		authRoutes.POST("/verify-email", func(c *gin.Context) { authHandler.VerifyEmail(c.Request.Context(), c) })
	}
}
