package repository

import (
	"gorm.io/gorm"

	"github.com/peter910820/kurohelper-db/models"
)

// 撈出日文漢字以及繁體中文字對應資料
func GetAllZhtwToJp(db *gorm.DB) ([]models.ZhtwToJp, error) {
	var results []models.ZhtwToJp

	err := db.Find(&results).Error

	if err != nil {
		return results, err
	}

	return results, nil
}
