package postgres

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConnOptions struct {
	User   string
	Pass   string
	Host   string
	Port   uint
	DBName string
	Schema string
}

func (o DBConnOptions) PostgresDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		o.Host, o.Port, o.User, o.Pass, o.DBName, o.Schema)
}

func NewPsqlGormConnection(opt DBConnOptions) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(opt.PostgresDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Enable logging for debugging
	})
}

// AutoMigrate runs the provided migrations
func AutoMigrate(db *gorm.DB, models ...interface{}) error {
	return db.AutoMigrate(models...)
}


func LoadDBConnOptionsFromEnv() DBConnOptions {
	port, _ := strconv.Atoi(getEnv("DB_PORT", "5432")) // Default to 5432 if not set

	return DBConnOptions{
		User:   getEnv("DB_USER", "maps_user"),
		Pass:   getEnv("DB_PASSWORD", "password123"),
		Host:   getEnv("DB_HOST", "localhost"),
		Port:   uint(port),
		DBName: getEnv("DB_NAME", "maps_db"),
		Schema: getEnv("DB_SCHEMA", "public"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
