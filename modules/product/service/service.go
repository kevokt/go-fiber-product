package service

import (
	"context"
	"go-fiber-modular/models"
)

type ProductService interface {
	CreateProduct(ctx context.Context, data ProductRequest) error
	ListByToko(ctx context.Context, tokoID int64) ([]*models.Product, error)
	GetProduct(ctx context.Context, productID int64) (*models.Product, error)
	UpdateProduct(ctx context.Context, data ProductRequest) error
	DeleteProduct(ctx context.Context, productID int64) error
}

type ProductRequest struct {
	ID          int64  `json:"id"`
	TokoID      int64  `json:"toko_id" validate:"required"`
	Product     string `json:"product" validate:"required"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity" validate:"required"`
}
