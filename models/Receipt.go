package models

import "time"

type Receipt struct {
	ID            uint      `gorm:"primary_key"`
	UserID        uint      `gorm:"not null"`
	TransactionID uint      `gorm:"not null"`
	Merchant      string    `gorm:"size:255; not null"`
	TotalAmount   float64   `gorm:"not null"`
	TaxAmount     float64
	Currency      string    `gorm:"size:3; not null; default:'USD'"`
	PurchasedAt   time.Time `gorm:"not null"`
	Notes         string    `gorm:"size:1024"`
	ImageURL      string    `gorm:"size:512"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
