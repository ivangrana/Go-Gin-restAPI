package controllers

import (
	"Phinance/database"
	DTOs "Phinance/dto"
	"Phinance/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllGoals(c *gin.Context) {
	var goals []models.Goals
	var goalsDTO []DTOs.GoalDTO
	resp := database.DB.Find(&goals, "user_id = ?", c.Param("id"))
	if resp.Error != nil {
		// Check if error is not due to no rows found
		if resp.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Goal not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": resp.Error.Error()})
		}
		return
	}
	//dto usage
	for _, goal := range goals {
		goalDTO := DTOs.GoalDTO{
			ID:     goal.ID,
			Amount: goal.Amount,
		}
		goalsDTO = append(goalsDTO, goalDTO)
	}
	c.JSON(http.StatusOK, goalsDTO)
}

func GetGoalById(c *gin.Context) {
	if c.Param("goal_id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "goal_id is required"})
		return
	}

	var goal models.Goals
	var goalDTO DTOs.GoalDTO

	resp := database.DB.First(&goal, c.Param("goal_id"))
	fmt.Println(resp.Error)
	if resp.Error != nil {
		// Check if error is not due to no rows found
		if resp.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Goal not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": resp.Error.Error()})
		}
		return
	} else {
		goalDTO = DTOs.GoalDTO{
			ID:     goal.ID,
			Amount: goal.Amount,
		}
		c.JSON(http.StatusOK, goalDTO)
	}

}

func CreateGoal(c *gin.Context) {
	var goal models.Goals
	var goalDTO DTOs.GoalCreateDTO

	//convert string to uint
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&goalDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	goal.UserID = uint(id)
	goal.Amount = goalDTO.Amount

	database.DB.Create(&goal)
	c.JSON(http.StatusCreated, gin.H{"message": "goal created"})
}

func DeleteGoal(c *gin.Context) {
	var goal models.Goals
	id := c.Param("goal_id")
	// database.DB.First(&goal, "id = ?", id)

	if err := database.DB.Delete(&goal, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Response": "Goal deleted"})
}
