package models

type Data struct {
	Symbol    string  `json:"symbol" db:"symbol"`
	Price     float64 `json:"price_24h" db:"price"`
	Volume    float64 `json:"volume_24h" db:"volume"`
	LastTrade float64 `json:"last_trade_price" db:"last_trade"`
}
