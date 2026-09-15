package services

import (
	"Phinance/database"
	"Phinance/dto"
	"Phinance/models"
	"strconv"
)

func GetAllTransactions(userID string) ([]dto.TransactionDTO, error) {
	var transactions []models.Transactions
	var transactionDTOs []dto.TransactionDTO

	resp := database.DB.Find(&transactions, "user_ID = ?", userID)
	if resp.Error != nil {
		return nil, resp.Error
	}

	for _, transaction := range transactions {
		transactionDTOs = append(transactionDTOs, dto.TransactionDTO{
			ID:          transaction.ID,
			CategoryID:  transaction.CategoryID,
			Value:       transaction.Value,
			Description: transaction.Description,
			Date:        transaction.Date,
		})
	}

	return transactionDTOs, nil
}

func GetTransactionById(transactionID string) (*dto.TransactionDTO, error) {
	var transaction models.Transactions

	resp := database.DB.First(&transaction, transactionID)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return &dto.TransactionDTO{
		ID:          transaction.ID,
		CategoryID:  transaction.CategoryID,
		Value:       transaction.Value,
		Description: transaction.Description,
		Date:        transaction.Date,
	}, nil
}

func CreateTransaction(userID string, transactionDTO dto.TransactionCreateDTO) error {
	id, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}

	transaction := models.Transactions{
		UserID:      uint(id),
		CategoryID:  transactionDTO.CategoryID,
		Value:       transactionDTO.Value,
		Description: transactionDTO.Description,
		Date:        transactionDTO.Date,
	}

	return database.DB.Create(&transaction).Error
}

func DeleteTransaction(transactionID string) error {
	return database.DB.Delete(&models.Transactions{}, transactionID).Error
}
