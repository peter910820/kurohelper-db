package repository

import (
	"gorm.io/gorm"

	kurohelperdb "github.com/peter910820/kurohelper-db"
	"github.com/peter910820/kurohelper-db/models"
)

// 查詢白名單
//
// 參數為guild(群組)跟dm(私訊)
func GetDiscordAllowListByKind(db *gorm.DB, kind string) ([]models.DiscordAllowList, error) {
	var results []models.DiscordAllowList

	if kind != "guild" && kind != "dm" {
		return nil, kurohelperdb.ErrParameterNotFound
	}

	err := db.Where("kind = ?", kind).Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}
