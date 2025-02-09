package repository

import (
	"github.bom/Alym62/backend/korp-stock-service/internal/domain"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) Create(product *domain.Product) error {
	return r.DB.Create(product).Error
}

func (r *ProductRepository) GetById(id uint) (*domain.Product, error) {
	var product domain.Product
	err := r.DB.First(&product, id).Error
	return &product, err
}

func (r *ProductRepository) Update(product *domain.Product) error {
	return r.DB.Save(product).Error
}

func (r *ProductRepository) DeleteById(id uint) error {
	return r.DB.Delete(&domain.Product{}, id).Error
}

func (r *ProductRepository) GetList() ([]domain.Product, error) {
	var products []domain.Product
	err := r.DB.Find(&products).Error
	return products, err
}
