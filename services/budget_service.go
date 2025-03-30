package services

import (
	"Phinance/database"
	"Phinance/dto"
	"Phinance/models"
	"strconv"

	"gorm.io/gorm"
)

type BudgetService struct {
	db *gorm.DB
}

func NewBudgetService() *BudgetService {
	return &BudgetService{db: database.DB}
}

func (s *BudgetService) GetAllBudgets(userID string) ([]dto.BudgetDTO, error) {
	var budgets []models.Budget
	var budgetDTOs []dto.BudgetDTO

	resp := s.db.Find(&budgets, "user_ID = ?", userID)
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

func (s *BudgetService) GetBudgetById(budgetID string) (*dto.BudgetDTO, error) {
	var budget models.Budget

	resp := s.db.First(&budget, budgetID)
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

func (s *BudgetService) CreateBudget(userID string, budgetDTO dto.BudgetCreateDTO) error {
	id, _ := strconv.Atoi(userID)
	budget := models.Budget{
		UserID:      uint(id),
		LimitValue:  budgetDTO.LimitValue,
		InitialDate: budgetDTO.InitialDate,
		FinalDate:   budgetDTO.FinalDate,
	}

	return s.db.Create(&budget).Error
}

func (s *BudgetService) DeleteBudget(budgetID string) error {
	return s.db.Delete(&models.Budget{}, budgetID).Error
}
