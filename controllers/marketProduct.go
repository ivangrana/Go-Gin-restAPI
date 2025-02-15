package controllers

import (
	"Phinance/database"
	dto "Phinance/dto"
	"Phinance/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllProducts(c *gin.Context) {
	var marketProducts []models.MarketProduct
	var marketProductDTOs []dto.MarketProductDTO
	resp := database.DB.Find(&marketProducts)
	if resp.Error != nil {
		if resp.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "MarketProduct not found"})
		}
	}
	for _, marketProduct := range marketProducts {
		marketProductDTO := dto.MarketProductDTO{
			ID:       marketProduct.ID,
			Name:     marketProduct.ProductName,
			Priority: marketProduct.Priority,
			Price:    float64(marketProduct.AveragePrice),
		}
		marketProductDTOs = append(marketProductDTOs, marketProductDTO)
	}

	c.JSON(http.StatusOK, marketProductDTOs)
}

func CreateMarketProduct(c *gin.Context) {
	var marketProduct models.MarketProduct
	var marketProductDTO dto.MarketProductCreateDTO

	if err := c.ShouldBindJSON(&marketProductDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	marketProduct.ProductName = marketProductDTO.ProductName
	marketProduct.Priority = marketProductDTO.Priority
	marketProduct.AveragePrice = float32(marketProductDTO.Price)

	database.DB.Save(&marketProduct)

	c.JSON(http.StatusOK, gin.H{"message": "MarketProduct created successfully"})
}

func GetMarketProductById(c *gin.Context) {
	var marketProduct models.MarketProduct
	var marketProductDTO dto.MarketProductDTO
	id := c.Param("market_product_id")
	database.DB.First(&marketProduct, "id = ?", id)

	marketProductDTO = dto.MarketProductDTO{
		ID:       marketProduct.ID,
		Name:     marketProduct.ProductName,
		Priority: marketProduct.Priority,
		Price:    float64(marketProduct.AveragePrice),
	}
	c.JSON(http.StatusOK, marketProductDTO)
}

func UpdateMarketProduct(c *gin.Context) {
	var marketProduct models.MarketProduct
	id := c.Param("market_product_id")
	database.DB.First(&marketProduct, "id = ?", id)
	if err := c.ShouldBindJSON(&marketProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Save(&marketProduct).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "MarketProduct updated successfully"})

}

func DeleteMarketProduct(c *gin.Context) {
	var marketProduct models.MarketProduct
	id := c.Param("product_id")
	if err := database.DB.Delete(&marketProduct, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "MarketProduct deleted successfully"})

}
