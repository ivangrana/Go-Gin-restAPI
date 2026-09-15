package controllers

import (
	dto "Phinance/dto"
	"Phinance/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllReceipts(c *gin.Context) {
	receipts, err := services.GetAllReceipts(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, receipts)
}

func GetReceiptById(c *gin.Context) {
	if c.Param("receipt_id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "receipt_id is required"})
		return
	}

	receipt, err := services.GetReceiptById(c.Param("receipt_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, receipt)
}

func CreateReceipt(c *gin.Context) {
	var receiptDTO dto.ReceiptCreateDTO
	if err := c.ShouldBindJSON(&receiptDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if c.Param("id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user id is required"})
		return
	}

	if err := services.CreateReceipt(c.Param("id"), receiptDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "receipt created"})
}

func UpdateReceipt(c *gin.Context) {
	var receiptDTO dto.ReceiptUpdateDTO
	if err := c.ShouldBindJSON(&receiptDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if c.Param("receipt_id") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "receipt_id is required"})
		return
	}

	if err := services.UpdateReceipt(c.Param("receipt_id"), receiptDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "receipt updated"})
}

func DeleteReceipt(c *gin.Context) {
	if err := services.DeleteReceipt(c.Param("receipt_id")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "receipt deleted"})
}
