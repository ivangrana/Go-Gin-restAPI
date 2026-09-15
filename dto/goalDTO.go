package dto

type GoalDTO struct {
	ID       uint    `json:"id"`
	Amount   float64 `json:"amount"`
	BudgetID uint    `json:"budget_id"`
}

type GoalCreateDTO struct {
	Amount   float64 `json:"amount"`
	BudgetID uint    `json:"budget_id"`
}

type GoalUpdateDTO struct {
	Amount   float64 `json:"amount"`
	BudgetID uint    `json:"budget_id"`
}
