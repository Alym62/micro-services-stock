package domain

import "gorm.io/gorm"

type Invoice struct {
	gorm.Model
	ID       uint
	Products []Product `gorm:"many2many:invoice_products;"`
	Status   string    `json:"status"`
}

type InvoiceProduct struct {
	gorm.Model
	ProductId uint `gorm:"primary_key"`
	InvoiceId uint `gorm:"primary_key"`
	Quantity  int
}
