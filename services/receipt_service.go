package services

import (
	"Phinance/database"
	"Phinance/dto"
	"Phinance/models"
	"strconv"

	"gorm.io/gorm"
)

func GetAllReceipts(userID string) ([]dto.ReceiptDTO, error) {
	var receipts []models.Receipt
	var receiptDTOs []dto.ReceiptDTO

	resp := database.DB.Find(&receipts, "user_id = ?", userID)
	if resp.Error != nil {
		return nil, resp.Error
	}

	for _, r := range receipts {
		receiptDTOs = append(receiptDTOs, dto.ReceiptDTO{
			ID:            r.ID,
			UserID:        r.UserID,
			TransactionID: r.TransactionID,
			Merchant:      r.Merchant,
			TotalAmount:   r.TotalAmount,
			TaxAmount:     r.TaxAmount,
			Currency:      r.Currency,
			PurchasedAt:   r.PurchasedAt,
			Notes:         r.Notes,
			ImageURL:      r.ImageURL,
			CreatedAt:     r.CreatedAt,
			UpdatedAt:     r.UpdatedAt,
		})
	}

	return receiptDTOs, nil
}

func GetReceiptById(receiptID string) (*dto.ReceiptDTO, error) {
	var r models.Receipt

	resp := database.DB.First(&r, receiptID)
	if resp.Error != nil {
		if resp.Error == gorm.ErrRecordNotFound {
			return nil, resp.Error
		}
		return nil, resp.Error
	}

	return &dto.ReceiptDTO{
		ID:            r.ID,
		UserID:        r.UserID,
		TransactionID: r.TransactionID,
		Merchant:      r.Merchant,
		TotalAmount:   r.TotalAmount,
		TaxAmount:     r.TaxAmount,
		Currency:      r.Currency,
		PurchasedAt:   r.PurchasedAt,
		Notes:         r.Notes,
		ImageURL:      r.ImageURL,
		CreatedAt:     r.CreatedAt,
		UpdatedAt:     r.UpdatedAt,
	}, nil
}

func CreateReceipt(userID string, receiptDTO dto.ReceiptCreateDTO) error {
	id, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}

	receipt := models.Receipt{
		UserID:        uint(id),
		TransactionID: receiptDTO.TransactionID,
		Merchant:      receiptDTO.Merchant,
		TotalAmount:   receiptDTO.TotalAmount,
		TaxAmount:     receiptDTO.TaxAmount,
		Currency:      receiptDTO.Currency,
		PurchasedAt:   receiptDTO.PurchasedAt,
		Notes:         receiptDTO.Notes,
		ImageURL:      receiptDTO.ImageURL,
	}

	if receipt.Currency == "" {
		receipt.Currency = "USD"
	}

	return database.DB.Create(&receipt).Error
}

func UpdateReceipt(receiptID string, receiptDTO dto.ReceiptUpdateDTO) error {
	id, err := strconv.Atoi(receiptID)
	if err != nil {
		return err
	}

	editReceipt := models.Receipt{
		ID:          uint(id),
		Merchant:    receiptDTO.Merchant,
		TotalAmount: receiptDTO.TotalAmount,
		TaxAmount:   receiptDTO.TaxAmount,
		Notes:       receiptDTO.Notes,
		ImageURL:    receiptDTO.ImageURL,
	}

	return database.DB.Updates(&editReceipt).Error
}

func DeleteReceipt(receiptID string) error {
	return database.DB.Delete(&models.Receipt{}, receiptID).Error
}
