package repository

import (
	"context"
	"go-fiber-modular/models"
)

type TokoRepository interface {
	CreateToko(ctx context.Context, toko *models.Toko) error
	GetMyToko(ctx context.Context, userID int64) ([]*models.Toko, error)
	GetTokoByID(ctx context.Context, tokoID int64) (*models.Toko, error)
	UpdateToko(ctx context.Context, toko *models.Toko) error
}
