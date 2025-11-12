package kurohelperdb

import (
	"time"

	"gorm.io/gorm/clause"
)

// 取得指定使用著遊玩資料
func GetUserData(userID string) ([]UserGameErogs, error) {
	var userGames []UserGameErogs

	err := Dbs.
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

	err := Dbs.Where("id = ?", userID).FirstOrCreate(&user, User{ID: userID, Name: userName}).Error
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
