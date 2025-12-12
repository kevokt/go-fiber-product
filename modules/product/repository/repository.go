package repository

import (
	"context"
	"go-fiber-modular/models"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *models.Product) error
	ListByToko(ctx context.Context, tokoID int64) ([]*models.Product, error)
	GetProduct(ctx context.Context, productID int64) (*models.Product, error)
	UpdateProduct(ctx context.Context, product *models.Product) error
	DeleteProduct(ctx context.Context, productID int64) error
}
