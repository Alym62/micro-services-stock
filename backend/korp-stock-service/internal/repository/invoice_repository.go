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

func (r *InvoiceRepository) GetById(id uint) (*domain.Invoice, error) {
	var invoice domain.Invoice
	err := r.DB.First(&invoice, id).Error
	return &invoice, err
}

func (r *InvoiceRepository) GetList() ([]domain.Invoice, error) {
	var invoices []domain.Invoice
	err := r.DB.Find(&invoices).Error
	return invoices, err
}

func (r *InvoiceRepository) Update(invoice *domain.Invoice) error {
	return r.DB.Save(invoice).Error
}
