package database

import (
	"log"

	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) {
	if err := DB.AutoMigrate(
	//&usersfcm.TutorDeviceToken{},
	); err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
}
