package kurohelperdb

import (
	"time"

	"gorm.io/gorm/clause"
)

// 取得指定使用者資料
func GetUser(userID string) (User, error) {
	var user User

	err := dbs.First(&user, "id = ?", userID).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// 取得所有使用者資料
func GetAllUser() ([]User, error) {
	var user []User

	err := dbs.Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// 取得特定使用者的單筆遊玩資料
func GetUserGameErogs(userID string, gameErogsID int) (UserGameErogs, error) {
	var gameRecord UserGameErogs

	err := dbs.First(&gameRecord, "user_id = ? AND game_erogs_id = ?", userID, gameErogsID).Error
	if err != nil {
		return gameRecord, err
	}

	return gameRecord, nil
}

// 取得特定使用者的全部遊玩資料
func GetUserGameErogsByUserID(userID string) ([]UserGameErogs, error) {
	var userGameErogs []UserGameErogs

	err := dbs.Where("user_id = ?", userID).Find(&userGameErogs).Error
	if err != nil {
		return userGameErogs, err
	}

	return userGameErogs, nil
}

// 取得指定使用著遊玩資料
func GetUserData(userID string) ([]UserGameErogs, error) {
	var userGames []UserGameErogs

	err := dbs.
		Preload("GameErogs").
		Preload("GameErogs.BrandErogs").
		Where("user_id = ?", userID).
		Order("COALESCE(completed_at, created_at) DESC").
		Find(&userGames).Error
	if err != nil {
		return userGames, err
	}

	return userGames, nil
}

// 確保指定的User存在，不存在就直接建立
func FindOrCreateUser(userID string, userName string) (User, error) {
	var user User

	err := dbs.Where("id = ?", userID).FirstOrCreate(&user, User{ID: userID, Name: userName}).Error
	if err != nil {
		return user, err
	}

	return user, nil
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
		return dbs.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "user_id"}, {Name: "game_erogs_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"has_played", "in_wish", "updated_at", "completed_at"}),
		}).Create(ug).Error
	}

	return dbs.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "game_erogs_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"has_played", "in_wish", "updated_at"}),
	}).Create(ug).Error
}
