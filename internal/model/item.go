package model

import "time"

// Item is model for item table.
type Item struct {
	ID         int       `gorm:"primary_key;type:serial" json:"id"`
	UserID     int       `gorm:"type:int" json:"user_id"`
	Name       string    `gorm:"type:varchar" json:"name"`
	TaxCode    int       `gorm:"type:int" json:"tax_code"`
	Type       string    `gorm:"-" json:"type"`
	Refundable bool      `gorm:"-" json:"refundable"`
	Price      float64   `gorm:"type:numeric" json:"price"`
	Tax        float64   `gorm:"-" json:"tax"`
	Amount     float64   `gorm:"-" json:"amount"`
	CreatedAt  time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp" json:"updated_at"`
}
