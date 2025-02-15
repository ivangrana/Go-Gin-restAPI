package controllers

import (
	"Phinance/database"
	DTOs "Phinance/dto"
	"Phinance/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllBudgets(c *gin.Context) {
	var budgets []models.Budget
	var budgetDTOs []DTOs.BudgetDTO

	resp := database.DB.Find(&budgets, "user_ID = ?", c.Param("id"))
	if resp.Error != nil {
		// Check if error is not due to no rows found
		if resp.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Budget not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": resp.Error.Error()})
		}
		return
	}
	for _, budget := range budgets {
		budgetDTO := DTOs.BudgetDTO{
			ID:          budget.ID,
			LimitValue:  budget.LimitValue,
			InitialDate: budget.InitialDate,
			FinalDate:   budget.FinalDate,
		}
		budgetDTOs = append(budgetDTOs, budgetDTO)
	}
	c.JSON(http.StatusOK, budgetDTOs)
}

func GetBudgetById(c *gin.Context) {
	if c.Param("budget_id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "budget_id is required"})
		return
	}
	var budget models.Budget
	var budgetDTO DTOs.BudgetDTO

	resp := database.DB.First(&budget, c.Param("budget_id"))
	if resp.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp.Error.Error()})
		return
	}
	budgetDTO = DTOs.BudgetDTO{
		ID:          budget.ID,
		LimitValue:  budget.LimitValue,
		InitialDate: budget.InitialDate,
		FinalDate:   budget.FinalDate,
	}
	c.JSON(http.StatusOK, budgetDTO)
}

func CreateBudget(c *gin.Context) {
	var budget models.Budget
	var budgetDTO DTOs.BudgetCreateDTO

	//convert string to uint
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&budgetDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	budget.UserID = uint(id)
	budget.LimitValue = budgetDTO.LimitValue

	database.DB.Create(&budget)

	c.JSON(http.StatusCreated, gin.H{"message": "budget created successfully"})
}

func DeleteBudget(c *gin.Context) {
	var budget models.Budget
	id := c.Param("budget_id")

	if err := database.DB.Delete(&budget, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Response": "budget deleted"})
}
