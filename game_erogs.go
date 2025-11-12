package kurohelperdb

// 確保指定的GameErogs存在，不存在就直接建立
func FindOrCreateGameErogs(gameID int, gameName string, brandErogsID int) (GameErogs, error) {
	var gameErogs GameErogs

	err := Dbs.Where("id = ?", gameID).FirstOrCreate(&gameErogs, GameErogs{ID: gameID, Name: gameName, BrandErogsID: brandErogsID}).Error
	if err != nil {
		return gameErogs, err
	}

	return gameErogs, nil
}
