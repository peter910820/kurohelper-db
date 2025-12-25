package kurohelperdb

import (
	"time"

	"gorm.io/gorm/clause"
)

// 取得特定使用者的單筆遊玩資料
func GetUserGameErogs(userID string, gameErogsID int) (UserGameErogs, error) {
	var gameRecord UserGameErogs

	err := Dbs.First(&gameRecord, "user_id = ? AND game_erogs_id = ?", userID, gameErogsID).Error
	if err != nil {
		return gameRecord, err
	}

	return gameRecord, nil
}

// 取得特定使用者的全部遊玩資料
func GetUserGameErogsByUserID(userID string) ([]UserGameErogs, error) {
	var userGameErogs []UserGameErogs

	err := Dbs.Where("user_id = ?", userID).Find(&userGameErogs).Error
	if err != nil {
		return userGameErogs, err
	}

	return userGameErogs, nil
}

// 刪除指定的UserGameErogs資料
func DeleteUserGameErogs(userID string, gameErogsID int) error {
	result := Dbs.Delete(
		&UserGameErogs{},
		"user_id = ? AND game_erogs_id = ?",
		userID, gameErogsID,
	)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrNoRowsAffected
	}
	return nil
}

// 建立或更新指定的UserGameErogs資料
func UpsertUserGameErogs(userID string, gameErogsID int, hasPlayed bool, inWish bool, completeDate time.Time) error {
	ug := UserGameErogs{
		UserID:      userID,
		GameErogsID: gameErogsID,
		HasPlayed:   hasPlayed,
		InWish:      inWish,
		UpdatedAt:   time.Now(),
	}

	if !completeDate.IsZero() {
		ug.CompletedAt = &completeDate
		return Dbs.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}, {Name: "game_erogs_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"has_played", "in_wish", "updated_at", "completed_at"}),
		}).Create(ug).Error
	}

	return Dbs.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "game_erogs_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"has_played", "in_wish", "updated_at"}),
	}).Create(ug).Error
}
