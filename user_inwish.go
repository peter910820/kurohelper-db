package kurohelperdb

func CreateUserInWish(userID string, gameErogsID int) error {
	userInWish := UserInWish{
		UserID:      userID,
		GameErogsID: gameErogsID,
	}

	if err := dbs.Create(&userInWish).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUserInWish(userID string, gameErogsID int) error {
	result := dbs.
		Where("user_id = ? AND game_erogs_id = ?", userID, gameErogsID).
		Delete(&UserInWish{}).Error

	return result
}
