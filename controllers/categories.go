package controllers

import (
	"Phinance/database"
	dto "Phinance/dto"
	"Phinance/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllCategories(c *gin.Context) {
	var categories []models.Categories
	var categoryDTOs []dto.CategoryDTO

	resp := database.DB.Find(&categories)
	if resp.Error != nil {
		if resp.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		}
	}
	for _, category := range categories {
		categoryDTO := dto.CategoryDTO{
			ID:   category.ID,
			Name: category.Name,
		}
		categoryDTOs = append(categoryDTOs, categoryDTO)
	}

	c.JSON(http.StatusOK, categoryDTOs)
}

func CreateCategory(c *gin.Context) {
	var category models.Categories

	resp := database.DB.Create(&category)
	if resp.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": resp.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category created successfully"})
}

func GetCategoryById(c *gin.Context) {
	var category models.Categories
	var categoryDTO dto.CategoryDTO
	id := c.Param("category_id")
	database.DB.First(&category, "id = ?", id)

	categoryDTO = dto.CategoryDTO{
		ID:   category.ID,
		Name: category.Name,
	}
	c.JSON(http.StatusOK, categoryDTO)
}

func UpdateCategory(c *gin.Context) {
	var category models.Categories
	id := c.Param("category_id")
	database.DB.First(&category, "id = ?", id)
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Save(&category).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})

}

func DeleteCategory(c *gin.Context) {
	var category models.Categories
	id := c.Param("category_id")
	if err := database.DB.Delete(&category, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Response": "Category deleted"})
}
