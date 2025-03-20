package main

import (
	"auth/cmd/server"
	"auth/pkg/configs"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

// @title Auth Service API
// @version 1.0
// @description This is the authentication service for user management. It provides endpoints for user registration, login, and password management.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@example.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}
	r := gin.Default()

	server.AuthServer(r, config)

	r.Run(":8080")
}
