package kurohelperdb

import "gorm.io/gorm"

// 確保指定的GameErogs存在，不存在就直接建立
func EnsureGameErogs(gameID int, gameName string, brandID int) (*GameErogs, error) {
	var game GameErogs
	if err := Dbs.Where("id = ?", gameID).
		FirstOrCreate(&game, GameErogs{ID: gameID, Name: gameName, BrandErogsID: brandID}).Error; err != nil {
		return nil, err
	}
	return &game, nil
}

// 確保指定的GameErogs存在，不存在就直接建立(Tx版本)
func EnsureGameErogsTx(tx *gorm.DB, gameID int, gameName string, brandID int) (*GameErogs, error) {
	var game GameErogs
	if err := tx.Where("id = ?", gameID).
		FirstOrCreate(&game, GameErogs{ID: gameID, Name: gameName, BrandErogsID: brandID}).Error; err != nil {
		return nil, err
	}
	return &game, nil
}
