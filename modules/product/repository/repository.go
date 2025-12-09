package repository

import (
	"context"

	"go-fiber-modular/models"
)

type ProductRepository interface {
	FindByID(ctx context.Context, id int64) (*models.Product, error)
	Create(ctx context.Context, product *models.Product) error
	Update(ctx context.Context, product *models.Product) error
	Delete(ctx context.Context, product *models.Product) error
}
