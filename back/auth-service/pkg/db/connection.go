package db

import (
	"auth/internal/domain"
	"auth/pkg/configs"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config *configs.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
		config.DBSSLMode,
	)

	log.Printf("Connecting to database with DSN: %s", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("postgres connection error: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get underlying *sql.DB: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	err = sqlDB.Ping()
	if err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	log.Println("Successfully connected to database")

	// Drop table if exists
	if db.Migrator().HasTable(&domain.Auth{}) {
		err = db.Migrator().DropTable(&domain.Auth{})
		if err != nil {
			log.Printf("Warning: failed to drop table: %v", err)
		}
	}

	// Auto migrate the schema
	err = db.AutoMigrate(&domain.Auth{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Successfully migrated database schema")

	return db
}
