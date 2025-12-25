package kurohelperdb

import "gorm.io/gorm"

func CreateUserInWish(userID string, gameErogsID int) error {
	userInWish := UserInWish{
		UserID:      userID,
		GameErogsID: gameErogsID,
	}

	if err := Dbs.Create(&userInWish).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUserInWish(userID string, gameErogsID int) error {
	result := Dbs.
		Where("user_id = ? AND game_erogs_id = ?", userID, gameErogsID).
		Delete(&UserInWish{}).Error

	return result
}

func CreateUserInWishTx(tx *gorm.DB, userID string, gameErogsID int) error {
	userInWish := UserInWish{
		UserID:      userID,
		GameErogsID: gameErogsID,
	}

	if err := tx.Create(&userInWish).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUserInWishTx(tx *gorm.DB, userID string, gameErogsID int) error {
	result := tx.
		Where("user_id = ? AND game_erogs_id = ?", userID, gameErogsID).
		Delete(&UserInWish{}).Error

	return result
}
