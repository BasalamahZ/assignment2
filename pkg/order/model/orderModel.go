package model

import (
	"time"
)

type Order struct {
	OrderId      int64     `json:"order_id" gorm:"primaryKey;UNIQUE"`
	CostumerName string    `json:"costumer_name"`
	OrderedAt    time.Time `json:"orderedAt" gorm:"autoCreateTime:true"`
	Items        []Item    `json:"items" gorm:"foreignKey:OrderId; constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
}

type Item struct {
	ItemId      int64     `json:"item_id" gorm:"primaryKey;UNIQUE"`
	ItemCode    string    `json:"item_code"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderId     int       `json:"order_id"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
}
