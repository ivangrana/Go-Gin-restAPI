package dto

type MarketProductDTO struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Priority string  `json:"priority"`
	Price    float64 `json:"price"`
}

type MarketProductCreateDTO struct {
	ProductName string  `json:"product_name"`
	Priority    string  `json:"priority"`
	Price       float64 `json:"price"`
}
