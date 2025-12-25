package kurohelperdb

import "gorm.io/gorm"

func EnsureUser(userID, userName string) (*User, error) {
	var user User
	if err := Dbs.Where("id = ?", userID).FirstOrCreate(&user, User{ID: userID, Name: userName}).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func EnsureUserTx(tx *gorm.DB, userID, userName string) (*User, error) {
	var user User
	if err := tx.Where("id = ?", userID).FirstOrCreate(&user, User{ID: userID, Name: userName}).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
