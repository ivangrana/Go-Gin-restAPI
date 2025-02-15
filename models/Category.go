package models

type Categories struct {
	ID           uint           `gorm:"primary_key"`
	Transactions []Transactions `gorm:"foreignKey:CategoryID"`
	Name         string
}
