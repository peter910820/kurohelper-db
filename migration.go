package kurohelperdb

import (
	"errors"

	"github.com/peter910820/kurohelper-db/models"
	"gorm.io/gorm"
)

// Migration
func Migration(db *gorm.DB) error {
	if db == nil {
		return errors.New("DB not initialized")
	}

	db.AutoMigrate(&models.ZhtwToJp{})
	db.AutoMigrate(&models.SeiyaCorrespond{})
	db.AutoMigrate(&models.WebAPIToken{})
	db.AutoMigrate(&models.DiscordAllowList{})
	db.AutoMigrate(
		&models.User{},
		&models.BrandErogs{},
		&models.GameErogs{},
		&models.UserGameErogs{},
	)

	return nil
}
