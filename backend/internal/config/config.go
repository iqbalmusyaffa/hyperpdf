package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config represents the complete application configuration
type Config struct {
	AppPort                  string `mapstructure:"APP_PORT"`
	AppEnv                   string `mapstructure:"APP_ENV"`
	DatabaseHost             string `mapstructure:"DATABASE_HOST"`
	DatabasePort             string `mapstructure:"DATABASE_PORT"`
	DatabaseUser             string `mapstructure:"DATABASE_USER"`
	DatabasePassword         string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseName             string `mapstructure:"DATABASE_NAME"`
	DatabaseSSLMode          string `mapstructure:"DATABASE_SSLMODE"`
	StorageUploadDir         string `mapstructure:"STORAGE_UPLOAD_DIR"`
	StorageOutputDir         string `mapstructure:"STORAGE_OUTPUT_DIR"`
	MaxFileSizeMB            int    `mapstructure:"MAX_FILE_SIZE_MB"`
	CompressorEngine         string `mapstructure:"COMPRESSOR_ENGINE"`
	GhostscriptBinary        string `mapstructure:"GHOSTSCRIPT_BINARY"`
	QPDFBinary               string `mapstructure:"QPDF_BINARY"`
	CompressTimeoutSeconds   int    `mapstructure:"COMPRESS_TIMEOUT_SECONDS"`
	RateLimitMax             int    `mapstructure:"RATE_LIMIT_MAX"`
	RateLimitDurationSeconds int    `mapstructure:"RATE_LIMIT_DURATION_SECONDS"`
}

// LoadConfig reads configuration from file or environment variables
func LoadConfig(configPath string) (*Config, error) {
	v := viper.New()

	// Set default values
	v.SetDefault("APP_PORT", "8080")
	v.SetDefault("APP_ENV", "development")
	v.SetDefault("DATABASE_HOST", "localhost")
	v.SetDefault("DATABASE_PORT", "5432")
	v.SetDefault("DATABASE_USER", "postgres")
	v.SetDefault("DATABASE_PASSWORD", "postgres")
	v.SetDefault("DATABASE_NAME", "pdf_tools")
	v.SetDefault("DATABASE_SSLMODE", "disable")
	v.SetDefault("STORAGE_UPLOAD_DIR", "./storage/uploads")
	v.SetDefault("STORAGE_OUTPUT_DIR", "./storage/outputs")
	v.SetDefault("MAX_FILE_SIZE_MB", 50)
	v.SetDefault("COMPRESSOR_ENGINE", "ghostscript")
	v.SetDefault("GHOSTSCRIPT_BINARY", "")
	v.SetDefault("QPDF_BINARY", "")
	v.SetDefault("COMPRESS_TIMEOUT_SECONDS", 120)
	v.SetDefault("RATE_LIMIT_MAX", 100)
	v.SetDefault("RATE_LIMIT_DURATION_SECONDS", 60)

	// Search config file
	if configPath != "" {
		v.SetConfigFile(configPath)
	} else {
		v.AddConfigPath(".")
		v.AddConfigPath("../")
		v.AddConfigPath("../../")
		v.SetConfigName(".env")
		v.SetConfigType("env")
	}

	// Read environment variables
	v.AutomaticEnv()

	// If a config file is found, read it in
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok && !os.IsNotExist(err) {
			// Non-critical if .env file is missing since env vars work
			_ = err
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal configuration: %w", err)
	}

	// Ensure directories exist
	if err := os.MkdirAll(filepath.Clean(cfg.StorageUploadDir), 0755); err != nil {
		return nil, fmt.Errorf("failed to create upload directory %s: %w", cfg.StorageUploadDir, err)
	}
	if err := os.MkdirAll(filepath.Clean(cfg.StorageOutputDir), 0755); err != nil {
		return nil, fmt.Errorf("failed to create output directory %s: %w", cfg.StorageOutputDir, err)
	}

	return &cfg, nil
}

// GetDSN returns PostgreSQL connection DSN string for GORM
func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DatabaseHost,
		c.DatabasePort,
		c.DatabaseUser,
		c.DatabasePassword,
		c.DatabaseName,
		c.DatabaseSSLMode,
	)
}

// GetURLDSN returns PostgreSQL URL format DSN string for golang-migrate
func (c *Config) GetURLDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DatabaseUser,
		c.DatabasePassword,
		c.DatabaseHost,
		c.DatabasePort,
		c.DatabaseName,
		c.DatabaseSSLMode,
	)
}
