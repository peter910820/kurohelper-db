package kurohelperdb

import "gorm.io/gorm"

// 確保指定的BrandErogs存在，不存在就直接建立
func EnsureBrandErogs(brandID int, brandName string) (*BrandErogs, error) {
	var brand BrandErogs
	if err := Dbs.Where("id = ?", brandID).FirstOrCreate(&brand, BrandErogs{ID: brandID, Name: brandName}).Error; err != nil {
		return nil, err
	}
	return &brand, nil
}

// 確保指定的BrandErogs存在，不存在就直接建立(Tx版本)
func EnsureBrandErogsTx(tx *gorm.DB, brandID int, brandName string) (*BrandErogs, error) {
	var brand BrandErogs
	if err := tx.Where("id = ?", brandID).FirstOrCreate(&brand, BrandErogs{ID: brandID, Name: brandName}).Error; err != nil {
		return nil, err
	}
	return &brand, nil
}
