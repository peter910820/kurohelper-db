package kurohelperdb

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

func CreateUserInWish(userID string, gameErogsID int) error {
	userInWish := UserInWish{
		UserID:      userID,
		GameErogsID: gameErogsID,
	}

	if err := Dbs.Create(&userInWish).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrUniqueViolation
		}
		return err
	}

	return nil
}

func DeleteUserInWish(userID string, gameErogsID int) error {
	err := Dbs.
		Where("user_id = ? AND game_erogs_id = ?", userID, gameErogsID).
		Delete(&UserInWish{}).Error

	return err
}

func CreateUserInWishTx(tx *gorm.DB, userID string, gameErogsID int) error {
	userInWish := UserInWish{
		UserID:      userID,
		GameErogsID: gameErogsID,
	}

	if err := tx.Create(&userInWish).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrUniqueViolation
		}
		return err
	}

	return nil
}

func DeleteUserInWishTx(tx *gorm.DB, userID string, gameErogsID int) error {
	err := tx.
		Where("user_id = ? AND game_erogs_id = ?", userID, gameErogsID).
		Delete(&UserInWish{}).Error

	return err
}
