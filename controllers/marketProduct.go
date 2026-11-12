package controllers

import (
	dto "Phinance/dto"
	"Phinance/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	marketProducts, err := services.GetAllMarketProducts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, marketProducts)
}

func CreateMarketProduct(c *gin.Context) {
	var marketProductDTO dto.MarketProductCreateDTO

	if err := c.ShouldBindJSON(&marketProductDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := services.CreateMarketProduct(marketProductDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "MarketProduct created successfully"})
}

func GetMarketProductById(c *gin.Context) {
	if c.Param("market_product_id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "market_product_id is required"})
		return
	}

	marketProduct, err := services.GetMarketProductById(c.Param("market_product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, marketProduct)
}

func UpdateMarketProduct(c *gin.Context) {
	var marketProductDTO dto.MarketProductUpdateDTO
	if err := c.ShouldBindJSON(&marketProductDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if c.Param("market_product_id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "market_product_id is required"})
		return
	}

	err := services.UpdateMarketProduct(c.Param("market_product_id"), marketProductDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "MarketProduct updated successfully"})
}

func DeleteMarketProduct(c *gin.Context) {
	id := c.Param("product_id")

	err := services.DeleteMarketProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "MarketProduct deleted successfully"})
}
