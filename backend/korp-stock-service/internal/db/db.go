package db

import (
	"fmt"

	"github.bom/Alym62/backend/korp-stock-service/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializerPostgreSQL() (*gorm.DB, error) {
	db, err := gorm.Open(
		postgres.Open("host=localhost user=postgres password=postgres dbname=stock port=5432"), &gorm.Config{},
	)
	if err != nil {
		fmt.Printf("PostgreSQL error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&domain.Product{}, &domain.Invoice{}, &domain.InvoiceProduct{})
	if err != nil {
		fmt.Printf("PostgreSQL automigrate error: %v", err)
		return nil, err
	}

	return db, nil
}
