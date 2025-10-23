package repository

import (
	"time"

	kurohelperdb "github.com/peter910820/kurohelper-db"
	"github.com/peter910820/kurohelper-db/models"
)

// 取出所有的web api token
func GetWebAPIToken() ([]models.WebAPIToken, error) {
	var tokens []models.WebAPIToken
	if err := kurohelperdb.Dbs.Find(&tokens).Error; err != nil {
		return nil, err
	}
	return tokens, nil
}

// expiresDuration為Token的有效時間，無期限expiresDuration傳0
func CreateWebAPIToken(id string, expiresDuration time.Duration) error {
	var expiresAt *time.Time

	if expiresDuration > 0 {
		t := time.Now().Add(expiresDuration)
		expiresAt = &t
	}

	token := &models.WebAPIToken{
		ID:        id,
		ExpiresAt: expiresAt,
	}

	if err := kurohelperdb.Dbs.Create(token).Error; err != nil {
		return err
	}

	return nil
}
