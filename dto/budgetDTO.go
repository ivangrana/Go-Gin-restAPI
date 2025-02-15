package dto

import "time"

type BudgetDTO struct {
	ID          uint      `json:"id"`
	LimitValue  float64   `json:"limit_value"`
	InitialDate time.Time `json:"initial_date"`
	FinalDate   time.Time `json:"final_date"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type BudgetCreateDTO struct {
	LimitValue  float64   `json:"limit_value"`
	InitialDate time.Time `json:"initial_date"`
	FinalDate   time.Time `json:"final_date"`
}
