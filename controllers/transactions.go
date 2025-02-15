package controllers

import (
	"Phinance/database"
	dto "Phinance/dto"
	"Phinance/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllTransactions(c *gin.Context) {
	var transactions []models.Transactions
	var transactionDTOs []dto.TransactionDTO
	resp := database.DB.Find(&transactions)
	if resp.Error != nil {
		if resp.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		}
	}
	for _, transaction := range transactions {
		transactionDTO := dto.TransactionDTO{
			ID:          transaction.ID,
			CategoryID:  transaction.CategoryID,
			Value:       transaction.Value,
			Description: transaction.Description,
			Date:        transaction.Date,
		}
		transactionDTOs = append(transactionDTOs, transactionDTO)
	}

	c.JSON(http.StatusOK, transactionDTOs)
}

func CreateTransaction(c *gin.Context) {
	var transaction models.Transactions
	var transactionDTO dto.TransactionCreateDTO

	if err := c.ShouldBindJSON(&transactionDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction.CategoryID = transactionDTO.CategoryID
	transaction.Value = transactionDTO.Value
	transaction.Description = transactionDTO.Description
	transaction.Date = transactionDTO.Date

	database.DB.Create(&transaction)

	c.JSON(http.StatusOK, gin.H{"message": "Transaction created successfully"})
}

func GetTransactionById(c *gin.Context) {
	var transaction models.Transactions
	var transactionDTO dto.TransactionDTO
	id := c.Param("transaction_id")
	database.DB.First(&transaction, "id = ?", id)

	transactionDTO = dto.TransactionDTO{
		ID:          transaction.ID,
		CategoryID:  transaction.CategoryID,
		Value:       transaction.Value,
		Description: transaction.Description,
		Date:        transaction.Date,
	}
	c.JSON(http.StatusOK, transactionDTO)
}

func DeleteTransaction(c *gin.Context) {
	var transaction models.Transactions
	id := c.Param("transaction_id")

	if err := database.DB.Delete(&transaction, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Response": "Transaction deleted"})
}
