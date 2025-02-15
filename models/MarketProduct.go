package models

type MarketProduct struct {
	ID           uint    `gorm:"primary_key"`
	ProductName  string  `json:"product_name"`
	AveragePrice float32 `gorm:"not null"`
	Priority     string  `json:"priority"`
}
