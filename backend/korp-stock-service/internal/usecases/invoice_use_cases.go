package usecases

import (
	"fmt"

	"github.bom/Alym62/backend/korp-stock-service/internal/domain"
	"github.bom/Alym62/backend/korp-stock-service/internal/queue"
	"github.bom/Alym62/backend/korp-stock-service/internal/repository"
)

type InvoiceUseCase struct {
	Repository  *repository.InvoiceRepository
	ProductRepo *repository.ProductRepository
	Publisher   *queue.RabbitMQPublisher
}

func NewInvoiceUseCase(repo *repository.InvoiceRepository, productRepo *repository.ProductRepository, publi *queue.RabbitMQPublisher) *InvoiceUseCase {
	return &InvoiceUseCase{Repository: repo, ProductRepo: productRepo, Publisher: publi}
}

func (pu *InvoiceUseCase) CreateInvoice(products []domain.Product, quantities []int) (*domain.Invoice, error) {
	for i, product := range products {
		if product.Quantity < quantities[i] {
			return nil, fmt.Errorf("estoque insuficiente para o produto %s", product.Name)
		}
	}

	invoice := domain.Invoice{
		Status: "Criada",
	}

	for i, product := range products {
		invoiceProduct := domain.InvoiceProduct{
			ProductId: product.ID,
			Quantity:  quantities[i],
		}
		invoice.Products = append(invoice.Products, product)

		product.Quantity -= quantities[i]
		err := pu.ProductRepo.Update(&product)
		if err != nil {
			fmt.Printf("Ops! Não foi possível atualizar esse produto: %v", err)
			return nil, err
		}

		err = pu.Repository.DB.Create(&invoiceProduct).Error
		if err != nil {
			fmt.Printf("Ops! Não foi possível salvar o produto e a nota fiscal: %v", err)
			return nil, err
		}
	}

	err := pu.Repository.Create(&invoice)
	if err != nil {
		fmt.Printf("Ops! Não foi possível criar a nota fiscal: %v", err)
		return nil, err
	}

	err = pu.Publisher.Publish(invoice)
	if err != nil {
		fmt.Printf("Ops! Ocorreu um erro ao tentar enviar para a fila: %v", err)
		return nil, err
	}

	return &invoice, nil
}

func (pu *InvoiceUseCase) GetByIdInvoice(id uint) (*domain.Invoice, error) {
	return pu.Repository.GetById(id)
}

func (pu *InvoiceUseCase) GetListInvoice() ([]domain.Invoice, error) {
	return pu.Repository.GetList()
}
