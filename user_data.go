package kurohelperdb

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
