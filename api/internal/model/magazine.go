package model

type Magazine struct {
	ID     int     `json:"id"`
	Number string  `json:"number"`
	Price  float64 `json:"price"`
}
