package controllers

import (
	dto "Phinance/dto"
	"Phinance/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

var budgetService *services.BudgetService

func GetAllBudgets(c *gin.Context) {
	budgets, err := budgetService.GetAllBudgets(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, budgets)
}

func GetBudgetById(c *gin.Context) {
	if c.Param("budget_id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "budget_id is required"})
		return
	}

	budget, err := budgetService.GetBudgetById(c.Param("budget_id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, budget)
}

func CreateBudget(c *gin.Context) {
	var budgetDTO dto.BudgetCreateDTO
	if err := c.ShouldBindJSON(&budgetDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id is required"})
		return
	}

	resp := budgetService.CreateBudget(c.Param("id"), budgetDTO)

	if resp != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func DeleteBudget(c *gin.Context) {
	id := c.Param("budget_id")

	err := budgetService.DeleteBudget(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "budget deleted"})

}
