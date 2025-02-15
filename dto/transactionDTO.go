package dto

import "time"

type TransactionDTO struct {
	ID          uint    `json:"id"`
	CategoryID  uint    `json:"category_id"`
	Value       float64 `json:"value"`
	Description string  `json:"description"`
	Date        time.Time
}

type TransactionCreateDTO struct {
	CategoryID  uint    `json:"category_id"`
	Value       float64 `json:"value"`
	Description string  `json:"description"`
	Date        time.Time
}
