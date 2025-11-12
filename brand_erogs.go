package kurohelperdb

// 確保指定的BrandErogs存在，不存在就直接建立
func FindOrCreateBrandErogs(brandID int, brandName string) (BrandErogs, error) {
	var brandErogs BrandErogs

	err := dbs.Where("id = ?", brandID).FirstOrCreate(&brandErogs, BrandErogs{ID: brandID, Name: brandName}).Error
	if err != nil {
		return brandErogs, err
	}

	return brandErogs, nil
}
