package kurohelperdb

import (
	"time"

	"gorm.io/gorm"
)

func CreateUserHasPlayed(userID string, gameErogsID int, completedAt *time.Time) error {
	userHasPlayed := UserHasPlayed{
		UserID:      userID,
		GameErogsID: gameErogsID,
		CompletedAt: completedAt,
	}

	if err := Dbs.Create(&userHasPlayed).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUserHasPlayed(userID string, gameErogsID int) error {
	result := Dbs.
		Where("user_id = ? AND game_erogs_id = ?", userID, gameErogsID).
		Delete(&UserHasPlayed{}).Error

	return result
}

func CreateUserHasPlayedTx(tx *gorm.DB, userID string, gameErogsID int, completedAt *time.Time) error {
	userHasPlayed := UserHasPlayed{
		UserID:      userID,
		GameErogsID: gameErogsID,
		CompletedAt: completedAt,
	}

	if err := Dbs.Create(&userHasPlayed).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUserHasPlayedTx(tx *gorm.DB, userID string, gameErogsID int) error {
	result := Dbs.
		Where("user_id = ? AND game_erogs_id = ?", userID, gameErogsID).
		Delete(&UserHasPlayed{}).Error

	return result
}
