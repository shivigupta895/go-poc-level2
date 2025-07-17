package models

import "time"

type Payment struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	OrderID   string    `json:"order_id"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
}
