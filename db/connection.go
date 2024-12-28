// Package db provides a database connection and helper functions to interact with the database.
package db

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/faulteh/nap-and-go/config"
)

// DB holds the global database connection.
var DB *gorm.DB

// Connect initializes the database connection.
func Connect() {
	cfg := config.LoadDBConfig()

	// Custom logger for GORM
	newLogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Info,   // Log level
			IgnoreRecordNotFoundError: true,          // Ignore record not found errors
			Colorful:                  true,          // Enable color in logs
		},
	)

	var err error
	DB, err = gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Println("Database connection established")
}

// GetDB returns the database instance.
func GetDB() *gorm.DB {
	if DB == nil {
		// Try to connect
		Connect()
	}
	return DB
}
