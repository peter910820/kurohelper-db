package repository

import (
	kurohelperdb "github.com/peter910820/kurohelper-db"
	"github.com/peter910820/kurohelper-db/models"
)

// 取得指定使用著遊玩資料
func GetUserData(userID string) ([]models.UserGameErogs, error) {
	var userGames []models.UserGameErogs

	err := kurohelperdb.Dbs.
		Preload("GameErogs").
		Preload("GameErogs.BrandErogs").
		Where("user_id = ?", userID).
		Find(&userGames).Error

	if err != nil {
		return userGames, err
	}

	return userGames, nil
}
