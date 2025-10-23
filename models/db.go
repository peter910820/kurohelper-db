package models

import "time"

type (
	ZhtwToJp struct {
		ZhTw      string    `gorm:"primaryKey;size:1"` // 繁體中文漢字
		Jp        string    `gorm:"size:1;not null"`   // 日文漢字
		CreatedAt time.Time `gorm:"autoCreateTime"`
		UpdatedAt time.Time `gorm:"autoUpdateTime"`
	}

	// 誠也對應表，專門針對極端狀況去對應
	SeiyaCorrespond struct {
		GameName  string    `gorm:"primaryKey"`
		SeiyaURL  string    `gorm:"not null"`
		CreatedAt time.Time `gorm:"autoCreateTime"`
		UpdatedAt time.Time `gorm:"autoUpdateTime"`
	}
)

type (
	User struct {
		ID        string    `gorm:"primaryKey" json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
		UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	}

	GameErogs struct {
		ID           int       `gorm:"primaryKey;autoIncrement:false" json:"id"`
		BrandErogsID int       `json:"brandErogsId"`
		Name         string    `gorm:"unique" json:"name"` // 遊戲名稱(批評空間)
		CreatedAt    time.Time `gorm:"autoCreateTime" json:"createdAt"`
		UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updatedAt"`

		BrandErogs *BrandErogs `gorm:"foreignKey:BrandErogsID;references:ID" json:"brandErogs,omitempty"` // 單向 preload
	}

	BrandErogs struct {
		ID        int       `gorm:"primaryKey;autoIncrement:false" json:"id"`
		Name      string    `gorm:"unique" json:"name"`
		CreatedAt time.Time `gorm:"autoCreateTime" json:"createdAt"`
		UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updatedAt"`
	}

	UserGameErogs struct {
		UserID      string     `gorm:"primaryKey" json:"userId"`
		GameErogsID int        `gorm:"primaryKey;autoIncrement:false" json:"gameErogsId"`
		HasPlayed   bool       `json:"hasPlayed"`
		InWish      bool       `json:"inWish"`
		CompletedAt *time.Time `json:"completedAt,omitempty"`
		CreatedAt   time.Time  `gorm:"autoCreateTime" json:"createdAt"`
		UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updatedAt"`

		GameErogs *GameErogs `gorm:"foreignKey:GameErogsID;references:ID" json:"gameErogs,omitempty"` // 單向 preload
	}
)

type DiscordAllowList struct {
	ID         string    `gorm:"primaryKey"`
	Kind       string    `gorm:"not null"`
	Permission int       `gorm:"not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

type WebAPIToken struct {
	ID        string `gorm:"primaryKey"`
	ExpiresAt *time.Time
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (ZhtwToJp) TableName() string {
	return "zhtw_to_jp"
}
