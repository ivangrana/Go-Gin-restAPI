package models

type Goals struct {
	ID     uint    `gorm:"primary_key"`
	UserID uint    `gorm:"not null"`
	Amount float64 `gorm:"not null"`
}
