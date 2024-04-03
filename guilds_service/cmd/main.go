package main

import (
	"log"

	"github.com/X3ne/ds_ms/guilds_service/config"
	"github.com/X3ne/ds_ms/guilds_service/internal/models"
	"github.com/X3ne/ds_ms/guilds_service/internal/server"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase(cfg *config.Config) (db *gorm.DB, err error) {
	log.Println("Connecting to the database...")

	db, err = gorm.Open(sqlite.Open(cfg.DB.Path), &gorm.Config{TranslateError: true})
	if err != nil {
		return nil, err
	}

	log.Println("Connected to the database")

	return db, err
}

func SyncDatabase(db *gorm.DB) {
	log.Println("Syncing database...")

	err := db.AutoMigrate(
		&models.Guild{},
		&models.GuildMember{},
		&models.GuildBan{},
	)
	if err != nil {
		log.Fatalf("Failed to sync database: %v", err)
	}

	log.Println("Database synced")
}

func CloseDatabase(db *gorm.DB) {
	log.Println("Closing database connection...")
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
	}

	err = sqlDB.Close()
	if err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
	}

	log.Println("Database connection closed")
}

func main() {
	cfg := config.NewConfig()

	db, err := InitDatabase(cfg)
	defer CloseDatabase(db)

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	SyncDatabase(db)

	server.LaunchServer(cfg, db)
}
