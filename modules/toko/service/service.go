package service

import (
	"context"
	"go-fiber-modular/models"
)

type TokoService interface {
	CreateToko(ctx context.Context, data TokoRequest) error
	GetMyToko(ctx context.Context, userID int64) ([]*models.Toko, error)
	UpdateToko(ctx context.Context, data TokoRequest) error
}

type TokoRequest struct {
	ID      int64  `json:"id"`
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	UserID  int64  `grom:"index" json:"user_id" validate:"required"`
}
