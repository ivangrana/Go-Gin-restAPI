package services

import (
	"Phinance/database"
	"Phinance/dto"
	"Phinance/models"
	"strconv"
)

func GetAllBudgets(userID string) ([]dto.BudgetDTO, error) {
	var budgets []models.Budget
	var budgetDTOs []dto.BudgetDTO

	resp := database.DB.Find(&budgets, "user_ID = ?", userID)
	if resp.Error != nil {
		return nil, resp.Error
	}

	for _, budget := range budgets {
		budgetDTOs = append(budgetDTOs, dto.BudgetDTO{
			ID:          budget.ID,
			LimitValue:  budget.LimitValue,
			InitialDate: budget.InitialDate,
			FinalDate:   budget.FinalDate,
		})
	}

	return budgetDTOs, nil
}

func GetBudgetById(budgetID string) (*dto.BudgetDTO, error) {
	var budget models.Budget

	resp := database.DB.First(&budget, budgetID)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return &dto.BudgetDTO{
		ID:          budget.ID,
		LimitValue:  budget.LimitValue,
		InitialDate: budget.InitialDate,
		FinalDate:   budget.FinalDate,
	}, nil
}

func CreateBudget(userID string, budgetDTO dto.BudgetCreateDTO) error {
	id, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}

	budget := models.Budget{
		UserID:      uint(id),
		LimitValue:  budgetDTO.LimitValue,
		InitialDate: budgetDTO.InitialDate,
		FinalDate:   budgetDTO.FinalDate,
	}

	return database.DB.Create(&budget).Error
}

func UpdateBudget(budgetID string, budgetDTO dto.BudgetUpdateDTO) error {
	id, err := strconv.Atoi(budgetID)
	if err != nil {
		return err
	}

	editBudget := models.Budget{
		ID:         uint(id),
		LimitValue: budgetDTO.LimitValue,
	}

	return database.DB.Updates(&editBudget).Error

}

func DeleteBudget(budgetID string) error {
	return database.DB.Delete(&models.Budget{}, budgetID).Error
}
