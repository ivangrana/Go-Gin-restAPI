package models

import "time"

type Transactions struct {
	ID          uint `gorm:"primary_key"`
	UserID      uint `gorm:"not null"`
	CategoryID  uint `gorm:"not null"`
	Value       float64
	Description string
	Date        time.Time
}
