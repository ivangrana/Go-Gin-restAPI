package models

type User struct {
	ID       uint           `json:"id" gorm:"primaryKey"`
	Name     string         `json:"name"`
	Password string         `json:"password"`
	Budget   []Budget       `gorm:"foreignKey:UserID"`
	Transac  []Transactions `gorm:"foreignKey:UserID"`
	Goals    []Goals        `gorm:"foreignKey:UserID"`
}
