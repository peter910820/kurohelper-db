package kurohelperdb

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func UpsertUserGameErogsTransaction(txInput UpsertUserGameErogsTXInput) error {
	var user User
	var gameErogs GameErogs
	var brandErogs BrandErogs

	err := Dbs.Transaction(func(tx *gorm.DB) error {
		// 1. 確保 User 存在
		if err := tx.Where("id = ?", txInput.UserID).FirstOrCreate(&user, User{ID: txInput.UserID, Name: txInput.UserName}).Error; err != nil {
			return err
		}

		// 2. 確保 Brand 存在
		if err := tx.Where("id = ?", txInput.ErogsBrandID).FirstOrCreate(&brandErogs, BrandErogs{ID: txInput.ErogsBrandID, Name: txInput.ErogsBrandName}).Error; err != nil {
			return err
		}

		// 3. 確保 Game 存在
		if err := tx.Where("id = ?", txInput.ErogsGameID).FirstOrCreate(&gameErogs, GameErogs{ID: txInput.ErogsGameID, Name: txInput.ErogsGamename, BrandErogsID: txInput.ErogsBrandID}).Error; err != nil {
			return err
		}

		ug := UserGameErogs{
			UserID:      txInput.UserID,
			GameErogsID: txInput.ErogsGameID,
			HasPlayed:   txInput.HasPlayed,
			InWish:      txInput.InWish,
			UpdatedAt:   time.Now(),
		}
		// 4. 建立 UserGame
		if txInput.CompleteDate.IsZero() {
			result := tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "user_id"}, {Name: "game_erogs_id"}},
				DoUpdates: clause.AssignmentColumns([]string{"has_played", "in_wish", "updated_at"}),
			}).Create(&ug)
			if result.Error != nil {
				return result.Error
			}
		} else {
			ug.CompletedAt = &txInput.CompleteDate
			result := tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "user_id"}, {Name: "game_erogs_id"}},
				DoUpdates: clause.AssignmentColumns([]string{"has_played", "in_wish", "updated_at", "completed_at"}),
			}).Create(&ug)
			if result.Error != nil {
				return result.Error
			}
		}

		return nil // commit
	})
	if err != nil {
		return err
	}

	return nil
}
