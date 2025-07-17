package models

import "time"

type Order struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Item      string    `json:"item"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
