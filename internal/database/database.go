package db

import (
	"fmt"
	"log"
	"log/slog"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/melkishengue/gotemplate/internal/models"
	"github.com/melkishengue/gotemplate/pkg/utils"
)

var Connection *gorm.DB

func InitDatabase(config *gorm.Config) {
	var err error

	gormDB, err := gorm.Open(postgres.Open(GetConnectionString()), config)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = gormDB.AutoMigrate(
		&models.City{},
	)

	if err != nil {
		log.Fatalf("failed to auto-migrate the database: %s", err.Error())
	}

	Connection = gormDB

	slog.Info("Successfully connected to PostgreSQL!")
}

func GetConnectionString() string {
	host := utils.GetEnvOrDie("DB_HOST")
	port := utils.GetEnvOrDie("DB_PORT")
	dbName := utils.GetEnvOrDie("DB_DATABASE")
	user := utils.GetEnvOrDie("DB_USER")
	password := utils.GetEnvOrDie("DB_PASSWORD")
	sslmode := utils.GetEnvOrDie("DB_SSL_MODE_REQUIRED")

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, sslmode)
}
