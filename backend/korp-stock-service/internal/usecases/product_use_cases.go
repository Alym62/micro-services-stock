package usecases

import (
	"fmt"

	"github.bom/Alym62/backend/korp-stock-service/internal/domain"
	"github.bom/Alym62/backend/korp-stock-service/internal/queue"
	"github.bom/Alym62/backend/korp-stock-service/internal/repository"
)

type ProductUseCase struct {
	Repository *repository.ProductRepository
	Publisher  *queue.RabbitMQPublisher
}

func NewProductUseCase(repo *repository.ProductRepository, publi *queue.RabbitMQPublisher) *ProductUseCase {
	return &ProductUseCase{Repository: repo, Publisher: publi}
}

func (pu *ProductUseCase) CreateProduct(name string, price float64, description string, quantity int) (*domain.Product, error) {
	product := &domain.Product{
		Name:        name,
		Price:       price,
		Description: description,
		Quantity:    quantity,
	}

	err := pu.Repository.Create(product)
	return product, err
}

func (pu *ProductUseCase) GetByIdProduct(id uint) (*domain.Product, error) {
	return pu.Repository.GetById(id)
}

func (pu *ProductUseCase) UpdateProduct(id uint, name string, price float64, description string, quantity int) (*domain.Product, error) {
	product, err := pu.Repository.GetById(id)
	if err != nil {
		fmt.Printf("Ops! Não existe esse produto na base de dados: %v", err)
		return nil, err
	}

	product.Name = name
	product.Price = price
	product.Description = description
	product.Quantity = quantity

	err = pu.Repository.Update(product)
	return product, err
}

func (pu *ProductUseCase) DeleteProduct(id uint) error {
	return pu.Repository.DeleteById(id)
}

func (pu *ProductUseCase) GetListProduct() ([]domain.Product, error) {
	return pu.Repository.GetList()
}

func (pu *ProductUseCase) RemoveProduct(id uint, quantity int) error {
	product, err := pu.Repository.GetById(id)
	if err != nil {
		fmt.Printf("Ops! Não existe esse produto na base de dados: %v", err)
		return err
	}

	if product.Quantity < quantity {
		fmt.Printf("Ops! Estoque de produto insuficiente: %v", err)
		return err
	}

	product.Quantity -= quantity
	return pu.Repository.Update(product)
}
