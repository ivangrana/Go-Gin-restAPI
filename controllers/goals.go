package controllers

import (
	dto "Phinance/dto"
	"Phinance/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllGoals(c *gin.Context) {
	goals, err := services.GetAllGoals(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, goals)
}

func GetGoalById(c *gin.Context) {
	if c.Param("goal_id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "goal_id is required"})
		return
	}

	goal, err := services.GetGoalById(c.Param("goal_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, goal)
}

func CreateGoal(c *gin.Context) {
	var goalDTO dto.GoalCreateDTO

	if err := c.ShouldBindJSON(&goalDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id is required"})
		return
	}

	if err := services.CreateGoal(c.Param("id"), goalDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "goal created"})
}

func UpdateGoal(c *gin.Context) {
	var goalDTO dto.GoalUpdateDTO

	if err := c.ShouldBindJSON(&goalDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if c.Param("goal_id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "goal_id is required"})
		return
	}

	if err := services.UpdateGoal(c.Param("goal_id"), goalDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "goal updated"})
}

func DeleteGoal(c *gin.Context) {
	id := c.Param("goal_id")

	if err := services.DeleteGoal(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "goal deleted"})
}
