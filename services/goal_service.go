package services

import (
	"Phinance/database"
	"Phinance/dto"
	"Phinance/models"
	"strconv"
)

func GetAllGoals(userID string) ([]dto.GoalDTO, error) {
	var goals []models.Goals
	var goalDTOs []dto.GoalDTO

	resp := database.DB.Find(&goals, "user_id = ?", userID)
	if resp.Error != nil {
		return nil, resp.Error
	}

	for _, goal := range goals {
		goalDTOs = append(goalDTOs, dto.GoalDTO{
			ID:       goal.ID,
			Amount:   goal.Amount,
			BudgetID: goal.BudgetID,
		})
	}

	return goalDTOs, nil
}

func GetGoalById(goalID string) (*dto.GoalDTO, error) {
	var goal models.Goals

	resp := database.DB.First(&goal, goalID)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return &dto.GoalDTO{
		ID:       goal.ID,
		Amount:   goal.Amount,
		BudgetID: goal.BudgetID,
	}, nil
}

func CreateGoal(userID string, goalDTO dto.GoalCreateDTO) error {
	id, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}

	goal := models.Goals{
		UserID:   uint(id),
		Amount:   goalDTO.Amount,
		BudgetID: goalDTO.BudgetID,
	}

	return database.DB.Create(&goal).Error
}

func UpdateGoal(goalID string, goalDTO dto.GoalUpdateDTO) error {
	id, err := strconv.Atoi(goalID)
	if err != nil {
		return err
	}

	editGoal := models.Goals{
		ID:       uint(id),
		Amount:   goalDTO.Amount,
		BudgetID: goalDTO.BudgetID,
	}

	return database.DB.Updates(&editGoal).Error
}

func DeleteGoal(goalID string) error {
	return database.DB.Delete(&models.Goals{}, goalID).Error
}
