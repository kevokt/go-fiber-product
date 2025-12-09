package service

import (
	"context"

	"go-fiber-modular/models"
)

type ProductService interface {
	FindByID(ctx context.Context, id int64) (*models.Product, error)
	Create(ctx context.Context, data CreateProductData) (*models.Product, error)
	Update(ctx context.Context, data UpdateProductData) (*models.Product, error)
	DeleteByID(ctx context.Context, id int64) (*models.Product, error)
}

type CreateProductData struct {
	Product     string
	Description string
	Quantity    int
}

type UpdateProductData struct {
	ID          int64
	Product     string
	Description string
	Quantity    int
}
