package repository

import (
	"context"
	"go-fiber-modular/models"
)

type AuthRepository interface {
	Register(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}
