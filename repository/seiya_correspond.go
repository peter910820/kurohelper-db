package repository

import (
	"github.com/peter910820/kurohelper-db/models"
	"gorm.io/gorm"
)

// 撈出誠也對應資料
func GetAllSeiyaCorrespond(db *gorm.DB) ([]models.SeiyaCorrespond, error) {
	var results []models.SeiyaCorrespond

	err := db.Find(&results).Error

	if err != nil {
		return results, err
	}

	return results, nil
}
