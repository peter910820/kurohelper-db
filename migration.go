package kurohelperdb

import (
	"errors"

	"github.com/peter910820/kurohelper-db/models"
)

// Migration
func Migration() error {
	if Dbs == nil {
		return errors.New("DB not initialized")
	}

	Dbs.AutoMigrate(&models.ZhtwToJp{})
	Dbs.AutoMigrate(&models.SeiyaCorrespond{})
	Dbs.AutoMigrate(
		&models.User{},
		&models.BrandErogs{},
		&models.GameErogs{},
		&models.UserGameErogs{},
	)

	return nil
}
