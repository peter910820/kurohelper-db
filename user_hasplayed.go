package kurohelperdb

import (
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

func SelectUserHasPlayed(userID string) ([]UserHasPlayed, error) {
	var hasPlayed []UserHasPlayed

	err := Dbs.
		// Preload("User").
		Preload("GameErogs").
		Preload("GameErogs.BrandErogs").
		Where("user_id = ?", userID).
		Order("COALESCE(completed_at, created_at) DESC").
		Find(&hasPlayed).Error

	if err != nil {
		return nil, err
	}
	return hasPlayed, nil
}

func CreateUserHasPlayed(userID string, gameErogsID int, completedAt *time.Time) error {
	userHasPlayed := UserHasPlayed{
		UserID:      userID,
		GameErogsID: gameErogsID,
		CompletedAt: completedAt,
	}

	if err := Dbs.Create(&userHasPlayed).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrUniqueViolation
		}
		return err
	}

	return nil
}

func DeleteUserHasPlayed(userID string, gameErogsID int) error {
	err := Dbs.
		Where("user_id = ? AND game_erogs_id = ?", userID, gameErogsID).
		Delete(&UserHasPlayed{}).Error

	return err
}

func CreateUserHasPlayedTx(tx *gorm.DB, userID string, gameErogsID int, completedAt *time.Time) error {
	userHasPlayed := UserHasPlayed{
		UserID:      userID,
		GameErogsID: gameErogsID,
		CompletedAt: completedAt,
	}

	if err := tx.Create(&userHasPlayed).Error; err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return ErrUniqueViolation
		}
		return err
	}

	return nil
}

func DeleteUserHasPlayedTx(tx *gorm.DB, userID string, gameErogsID int) error {
	err := Dbs.
		Where("user_id = ? AND game_erogs_id = ?", userID, gameErogsID).
		Delete(&UserHasPlayed{}).Error

	return err
}
