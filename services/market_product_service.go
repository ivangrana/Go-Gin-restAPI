package services

import (
	"Phinance/database"
	"Phinance/dto"
	"Phinance/models"
	"strconv"
)

func GetAllMarketProducts() ([]dto.MarketProductDTO, error) {
	var marketProducts []models.MarketProduct
	var marketProductDTOs []dto.MarketProductDTO

	resp := database.DB.Find(&marketProducts)
	if resp.Error != nil {
		return nil, resp.Error
	}

	for _, marketProduct := range marketProducts {
		marketProductDTOs = append(marketProductDTOs, dto.MarketProductDTO{
			ID:       marketProduct.ID,
			Name:     marketProduct.ProductName,
			Priority: marketProduct.Priority,
			Price:    float64(marketProduct.AveragePrice),
		})
	}

	return marketProductDTOs, nil
}

func GetMarketProductById(marketProductID string) (*dto.MarketProductDTO, error) {
	var marketProduct models.MarketProduct

	resp := database.DB.First(&marketProduct, "id = ?", marketProductID)
	if resp.Error != nil {
		return nil, resp.Error
	}

	return &dto.MarketProductDTO{
		ID:       marketProduct.ID,
		Name:     marketProduct.ProductName,
		Priority: marketProduct.Priority,
		Price:    float64(marketProduct.AveragePrice),
	}, nil
}

func CreateMarketProduct(marketProductDTO dto.MarketProductCreateDTO) error {
	marketProduct := models.MarketProduct{
		ProductName:  marketProductDTO.ProductName,
		AveragePrice: float32(marketProductDTO.Price),
		Priority:     marketProductDTO.Priority,
	}

	return database.DB.Create(&marketProduct).Error
}

func UpdateMarketProduct(marketProductID string, marketProductDTO dto.MarketProductUpdateDTO) error {
	id, err := strconv.Atoi(marketProductID)
	if err != nil {
		return err
	}

	editMarketProduct := models.MarketProduct{
		ID:           uint(id),
		ProductName:  marketProductDTO.ProductName,
		AveragePrice: float32(marketProductDTO.Price),
		Priority:     marketProductDTO.Priority,
	}

	return database.DB.Updates(&editMarketProduct).Error
}

func DeleteMarketProduct(marketProductID string) error {
	return database.DB.Delete(&models.MarketProduct{}, "id = ?", marketProductID).Error
}
