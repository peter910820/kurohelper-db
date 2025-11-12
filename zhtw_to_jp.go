package kurohelperdb

// 撈出日文漢字以及繁體中文字對應資料
func GetAllZhtwToJp() ([]ZhtwToJp, error) {
	var results []ZhtwToJp

	err := dbs.Find(&results).Error

	if err != nil {
		return results, err
	}

	return results, nil
}
