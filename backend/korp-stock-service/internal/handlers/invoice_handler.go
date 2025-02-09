package handlers

import (
	"net/http"

	"github.bom/Alym62/backend/korp-stock-service/internal/domain"
	"github.bom/Alym62/backend/korp-stock-service/internal/usecases"
	"github.com/gin-gonic/gin"
)

type InvoiceHanlder struct {
	UseCase *usecases.InvoiceUseCase
}

func NewInvoiceHanlder(uc *usecases.InvoiceUseCase) *InvoiceHanlder {
	return &InvoiceHanlder{UseCase: uc}
}

func (h *InvoiceHanlder) CreateInvoice(c *gin.Context) {
	var request struct {
		Products   []domain.Product `json:"products"`
		Quantities []int            `json:"quantities"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	invoice, err := h.UseCase.CreateInvoice(request.Products, request.Quantities)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, invoice)
}
