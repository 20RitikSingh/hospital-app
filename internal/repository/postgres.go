package repository

import (
	"github.com/20ritiksingh/hospital-app/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.DSN
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
