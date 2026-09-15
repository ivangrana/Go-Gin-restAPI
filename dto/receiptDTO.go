package dto

import "time"

type ReceiptDTO struct {
	ID            uint      `json:"id"`
	UserID        uint      `json:"user_id"`
	TransactionID uint      `json:"transaction_id"`
	Merchant      string    `json:"merchant"`
	TotalAmount   float64   `json:"total_amount"`
	TaxAmount     float64   `json:"tax_amount"`
	Currency      string    `json:"currency"`
	PurchasedAt   time.Time `json:"purchased_at"`
	Notes         string    `json:"notes"`
	ImageURL      string    `json:"image_url"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type ReceiptCreateDTO struct {
	TransactionID uint      `json:"transaction_id" binding:"required"`
	Merchant      string    `json:"merchant"        binding:"required"`
	TotalAmount   float64   `json:"total_amount"    binding:"required"`
	TaxAmount     float64   `json:"tax_amount"`
	Currency      string    `json:"currency"`
	PurchasedAt   time.Time `json:"purchased_at"    binding:"required"`
	Notes         string    `json:"notes"`
	ImageURL      string    `json:"image_url"`
}

type ReceiptUpdateDTO struct {
	Merchant    string  `json:"merchant"`
	TotalAmount float64 `json:"total_amount"`
	TaxAmount   float64 `json:"tax_amount"`
	Notes       string  `json:"notes"`
	ImageURL    string  `json:"image_url"`
}
