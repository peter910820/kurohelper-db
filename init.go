package kurohelperdb

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/peter910820/kurohelper-db/models"
	"github.com/sirupsen/logrus"
)

var (
	Dbs *gorm.DB
)

func InitDsn(config models.Config) error {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DBOwner,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)

	// get connect db variable
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Errorf("連接資料庫失敗: %v", err)
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Errorf("無法取得 sql.DB: %v", err)
		return err
	}

	sqlDB.SetMaxOpenConns(30)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	Dbs = db
	return nil
}
