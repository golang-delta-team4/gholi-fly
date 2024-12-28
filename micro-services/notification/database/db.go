package database

import (
	"fmt"
	"notification-nats/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewConnection(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Database,
	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
