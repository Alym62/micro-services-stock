package domain

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
}
