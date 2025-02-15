package dto

type GoalDTO struct {
	ID     uint    `json:"id"`
	Amount float64 `json:"amount"`
}

type GoalCreateDTO struct {
	Amount float64 `json:"amount"`
}

type GoalUpdateDTO struct {
	Amount float64 `json:"amount"`
}
