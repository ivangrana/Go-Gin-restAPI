package controllers

import (
	"Phinance/database"
	dto "Phinance/dto"
	"Phinance/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Get all users
func GetAllUsers(c *gin.Context) {
	var users []models.User
	var userDTOs []dto.UserDTO
	resp := database.DB.Find(&users)
	if resp.Error != nil {
		if resp.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Users not found"})
		}
	}
	for _, user := range users {
		userDTO := dto.UserDTO{
			ID:       user.ID,
			Name:     user.Name,
			Password: user.Password,
		}
		userDTOs = append(userDTOs, userDTO)
	}

	c.JSON(http.StatusOK, userDTOs)
}

// Create a new user
func CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
}

// Get user by ID
func GetUserByID(c *gin.Context) {
	var user models.User
	var userDTO dto.UserDTO
	id := c.Param("id")
	database.DB.Find(&user, "id = ?", id)

	userDTO = dto.UserDTO{
		ID:       user.ID,
		Name:     user.Name,
		Password: user.Password,
	}
	c.JSON(http.StatusOK, userDTO)
}

// Update user
func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	database.DB.First(&user, "id = ?", id)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})

}

// Delete user
func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	database.DB.First(&user, "id = ?", id)
	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})

}
