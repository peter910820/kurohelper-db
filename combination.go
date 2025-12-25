package kurohelperdb

func BrandCount(userID string, hasPlayed bool, inWish bool) ([]BrandCountOuput, error) {
	var brandData []BrandCountOuput
	err := Dbs.
		Table("user_game_erogs AS uge").
		Select("b.id AS brand_id, b.name AS brand_name, COUNT(*) AS count").
		Joins("JOIN game_erogs AS g ON uge.game_erogs_id = g.id").
		Joins("JOIN brand_erogs AS b ON g.brand_erogs_id = b.id").
		Where("uge.user_id = ? AND uge.has_played = ?", userID, hasPlayed).
		Where("uge.user_id = ? AND uge.in_wish = ?", userID, inWish).
		Group("b.id, b.name").
		Order("count DESC").
		Scan(&brandData).Error
	if err != nil {
		return brandData, err
	}
	return brandData, nil
}
