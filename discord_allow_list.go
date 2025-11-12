package kurohelperdb

// 查詢白名單
//
// 參數為guild(群組)跟dm(私訊)
func GetDiscordAllowListByKind(kind string) ([]DiscordAllowList, error) {
	var results []DiscordAllowList

	if kind != "guild" && kind != "dm" {
		return nil, ErrParameterNotFound
	}

	err := dbs.Where("kind = ?", kind).Find(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}
