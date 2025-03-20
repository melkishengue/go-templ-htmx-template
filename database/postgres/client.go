package postgres

import (
	"fmt"
	"myapp/config"

	"database/sql"

	"github.com/gofiber/fiber/v2/log"
	_ "github.com/lib/pq"
)

var ref *sql.DB

func InitDatabase() *sql.DB {
	cfg := config.GetConfig()

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s %s", cfg.Database.Host, cfg.Database.User, cfg.Database.Database, cfg.Database.Password, cfg.Database.Port, cfg.Database.AdditionalParams)
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Unable to create connection pool: %v", err)
	}

	log.Info("Connection to database successful (" + cfg.Database.Host + ") 🚀 ")

	return db
}

func GetDB() *sql.DB {
	if ref == nil {
		ref = InitDatabase()
	}

	return ref
}
