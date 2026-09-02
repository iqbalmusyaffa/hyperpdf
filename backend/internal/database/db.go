package database

import (
	"errors"
	"fmt"
	"time"

	"cobagolang/backend/internal/config"
	"cobagolang/backend/internal/model"
	"cobagolang/backend/internal/utils"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// InitDB initializes PostgreSQL connection using GORM
func InitDB(cfg *config.Config) (*gorm.DB, error) {
	log := utils.GetLogger()

	gormLoggerLevel := logger.Warn
	if cfg.AppEnv == "development" {
		gormLoggerLevel = logger.Info
	}

	db, err := gorm.Open(postgres.Open(cfg.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(gormLoggerLevel),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Info("Connected to PostgreSQL successfully")
	return db, nil
}

// RunMigrations runs database migrations using golang-migrate
func RunMigrations(cfg *config.Config, migrationsPath string) error {
	log := utils.GetLogger()

	if migrationsPath == "" {
		migrationsPath = "file://migrations"
	} else {
		migrationsPath = "file://" + migrationsPath
	}

	m, err := migrate.New(migrationsPath, cfg.GetURLDSN())
	if err != nil {
		return fmt.Errorf("failed to initialize migration instance: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Info("Database migrations: no new changes")
			return nil
		}
		return fmt.Errorf("migration up failed: %w", err)
	}

	log.Info("Database migrations applied successfully")
	return nil
}

// AutoMigrate fallback for GORM models
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.PDFJob{},
	)
}
