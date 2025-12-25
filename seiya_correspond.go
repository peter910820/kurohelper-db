package kurohelperdb

// 撈出誠也對應資料
func GetAllSeiyaCorrespond() ([]SeiyaCorrespond, error) {
	var results []SeiyaCorrespond

	err := Dbs.Find(&results).Error

	if err != nil {
		return results, err
	}

	return results, nil
}
