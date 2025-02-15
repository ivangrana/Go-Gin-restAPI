package models

import (
	"time"
)

type Budget struct {
	ID          uint `gorm:"primary_key"`
	UserID      uint `gorm:"not null"`
	LimitValue  float64
	InitialDate time.Time
	FinalDate   time.Time
	UpdatedAt   time.Time
}
