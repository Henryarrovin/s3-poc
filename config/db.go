package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=127.0.0.1 user=postgres password=postgres dbname=s3demo port=5432 sslmode=disable"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Millisecond * 200, // log queries slower than 200ms
			LogLevel:      logger.Info,            // Log all SQL queries
			Colorful:      true,                   // Enable colored logs
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}
	return db, nil
}
