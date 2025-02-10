package main

import (
	"fmt"

	"github.bom/Alym62/backend/korp-stock-service/internal/db"
	"github.bom/Alym62/backend/korp-stock-service/internal/handlers"
	"github.bom/Alym62/backend/korp-stock-service/internal/queue"
	"github.bom/Alym62/backend/korp-stock-service/internal/repository"
	"github.bom/Alym62/backend/korp-stock-service/internal/usecases"
	"github.bom/Alym62/backend/korp-stock-service/pkg"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := db.InitializerPostgreSQL()
	if err != nil {
		fmt.Printf("Erro ao conectar com o PostgreSQL: %v", err)
	}

	rabbitUrl := "amqp://guest:guest@localhost:5672/"
	queueName := "products.v1.invoice-event"
	exhangeName := "products.v1.product"
	publisher, err := queue.NewRabbitMQPublisher(rabbitUrl, exhangeName, queueName)
	if err != nil {
		panic(err)
	}
	defer publisher.Close()

	productRepo := repository.NewProductRepository(db)
	productUseCase := usecases.NewProductUseCase(productRepo, publisher)
	productHandler := handlers.NewProductHanlder(productUseCase)

	r := gin.Default()
	r.Use(pkg.CORSMiddlewares())
	v1Product := r.Group("/api/v1/product")
	v1Product.GET("/list", productHandler.GetListProduct)
	v1Product.GET("/:id", productHandler.GetProduct)
	v1Product.POST("/create", productHandler.CreateProduct)
	v1Product.PUT("/update/:id", productHandler.UpdateProduct)
	v1Product.DELETE("/delete/:id", productHandler.DeleteProduct)

	invoiceRepo := repository.NewInvoiceRepository(db)
	invoiceUseCase := usecases.NewInvoiceUseCase(invoiceRepo, productRepo, publisher)
	invoiceHandler := handlers.NewInvoiceHanlder(invoiceUseCase)

	v1Invoice := r.Group("/api/v1/invoice")
	v1Invoice.Use(pkg.CORSMiddlewares())
	v1Invoice.POST("/create", invoiceHandler.CreateInvoice)
	v1Invoice.GET("/:id", invoiceHandler.GetInvoice)
	v1Invoice.GET("/list", invoiceHandler.GetListInvoice)

	r.Run(":8080")
}
