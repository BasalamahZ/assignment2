package dto

import (
	"time"
)

type OrderRequest struct {
	CostumerName string    `json:"costumer_name"`
	OrderedAt    time.Time `json:"orderedAt"`
	Item         []ItemRequest
}

type ItemRequest struct {
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type OrderResponse struct {
	OrderId      int64          `json:"order_id" gorm:"primaryKey;UNIQUE"`
	CostumerName string         `json:"costumer_name"`
	OrderedAt    time.Time      `json:"orderedAt"`
	Items        []ItemResponse `json:"items"`
	CreatedAt    time.Time      `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt    time.Time      `json:"updated_at" gorm:"autoUpdateTime:true"`
}

type ItemResponse struct {
	ItemId      int64     `json:"item_id" gorm:"primaryKey;UNIQUE"`
	ItemCode    string    `json:"item_code"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	OrderId     int       `json:"order_id"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime:true"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime:true"`
}
