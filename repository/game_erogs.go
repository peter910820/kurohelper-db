package repository

import (
	"github.com/peter910820/kurohelper-db/models"
	"gorm.io/gorm"
)

// 確保指定的GameErogs存在，不存在就直接建立
func FindOrCreateGameErogs(db *gorm.DB, gameID int, gameName string, brandErogsID int) (models.GameErogs, error) {
	var gameErogs models.GameErogs

	err := db.Where("id = ?", gameID).FirstOrCreate(&gameErogs, models.GameErogs{ID: gameID, Name: gameName, BrandErogsID: brandErogsID}).Error
	if err != nil {
		return gameErogs, err
	}

	return gameErogs, nil
}
