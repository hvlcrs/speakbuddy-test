package db

import (
	"fmt"
	"log"
	"speakbuddy/models"
	"speakbuddy/pkg/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Conn *gorm.DB = &gorm.DB{}

func Init() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		configs.PostgresSetting.Host,
		configs.PostgresSetting.User,
		configs.PostgresSetting.Password,
		configs.PostgresSetting.DBName,
		configs.PostgresSetting.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("[FAIL] Failed to open database connection: %v", err)
		return
	}

	// Populate database table schema using object models
	// In production, database migration can be handled by a separate tool
	// providing more robust features like versioning, rollback, etc.
	if err := db.AutoMigrate(&models.Audio{}); err != nil {
		log.Fatalf("[FAIL] Database migration failed: %v", err)
		return
	}

	Conn = db
}
