package models

type Price struct {
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Decimals int     `json:"decimals"`
}

type Item struct {
	Id           string `json:"id"`
	Title        string `json:"title"`
	Price        Price  `json:"price"`
	Picture      string `json:"picture"`
	Condition    string `json:"condition"`
	FreeShipping bool   `json:"free_shipping"`
	SoldQuantity int    `json:"sold_quantity"`
	Description  string `json:"description"`
}
