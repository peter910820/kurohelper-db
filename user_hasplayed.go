package kurohelperdb

import "time"

func CreateUserHasPlayed(userID string, gameErogsID int, completedAt *time.Time) error {
	userHasPlayed := UserHasPlayed{
		UserID:      userID,
		GameErogsID: gameErogsID,
		CompletedAt: completedAt,
	}

	if err := dbs.Create(&userHasPlayed).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUserHasPlayed(userID string, gameErogsID int) error {
	result := dbs.
		Where("user_id = ? AND game_erogs_id = ?", userID, gameErogsID).
		Delete(&UserHasPlayed{}).Error

	return result
}
