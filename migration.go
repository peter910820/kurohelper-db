package kurohelperdb

import (
	"gorm.io/gorm"

	"github.com/peter910820/kurohelper-db/models"
)

func Migration(db *gorm.DB) {
	db.AutoMigrate(&models.ZhtwToJp{})
	db.AutoMigrate(&models.SeiyaCorrespond{})
	db.AutoMigrate(
		&models.User{},
		&models.BrandErogs{},
		&models.GameErogs{},
		&models.UserGameErogs{},
	)
}
