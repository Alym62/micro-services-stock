package pkg

import (
	"encoding/json"
	"fmt"

	"github.com/Alym62/backend/korp-billing-service/internal/events"
	"github.com/signintech/gopdf"
)

func GeneratePDF(body []byte) error {
	var invoice events.InvoiceCreated
	err := json.Unmarshal(body, &invoice)
	if err != nil {
		return fmt.Errorf("erro ao desserializar JSON: %v", err)
	}

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()

	err = pdf.AddTTFFont("arial", "../assets/fonts/arial.ttf")
	if err != nil {
		return fmt.Errorf("erro ao carregar fonte: %v", err)
	}

	err = pdf.SetFont("arial", "", 14)
	if err != nil {
		return fmt.Errorf("erro ao definir fonte: %v", err)
	}

	pdf.Cell(nil, fmt.Sprintf("Nota Fiscal - ID: %d", invoice.ID))
	pdf.Br(20)

	pdf.Cell(nil, "Produtos:")
	pdf.Br(15)

	for _, p := range invoice.Products {
		total := p.Price * float64(p.Quantity)
		pdf.Cell(nil, fmt.Sprintf("%d x %s - R$%.2f (Total: R$%.2f)", p.Quantity, p.Name, p.Price, total))
		pdf.Br(12)
		pdf.Cell(nil, fmt.Sprintf("Descrição: %s", p.Description))
		pdf.Br(15)
	}

	pdf.Br(10)
	pdf.Cell(nil, fmt.Sprintf("Total: R$%.2f", calculateTotal(invoice.Products)))

	filePath := fmt.Sprintf("../pdf/invoice_%d.pdf", invoice.ID)
	err = pdf.WritePdf(filePath)
	if err != nil {
		return fmt.Errorf("erro ao gerar PDF: %v", err)
	}

	fmt.Println("Nota Fiscal gerada com sucesso:", filePath)

	return nil
}

func calculateTotal(products []events.Product) float64 {
	var total float64
	for _, p := range products {
		total += p.Price * float64(p.Quantity)
	}
	return total
}
