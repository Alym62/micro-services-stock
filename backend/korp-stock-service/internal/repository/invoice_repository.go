package repository

import (
	"github.bom/Alym62/backend/korp-stock-service/internal/domain"
	"gorm.io/gorm"
)

type InvoiceRepository struct {
	DB *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) *InvoiceRepository {
	return &InvoiceRepository{DB: db}
}

func (r *InvoiceRepository) Create(invoice *domain.Invoice) error {
	return r.DB.Create(invoice).Error
}
