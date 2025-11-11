package repository

import (
	"github.com/peter910820/kurohelper-db/models"
	"gorm.io/gorm"
)

// 確保指定的BrandErogs存在，不存在就直接建立
func FindOrCreateBrandErogs(db *gorm.DB, brandID int, brandName string) (models.BrandErogs, error) {
	var brandErogs models.BrandErogs

	err := db.Where("id = ?", brandID).FirstOrCreate(&brandErogs, models.BrandErogs{ID: brandID, Name: brandName}).Error
	if err != nil {
		return brandErogs, err
	}

	return brandErogs, nil
}
