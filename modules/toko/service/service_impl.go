package service

import (
	"context"
	"errors"
	"go-fiber-modular/models"
	"go-fiber-modular/modules/toko/repository"
)

type service struct {
	repo repository.TokoRepository
}

func NewTokoService(repo repository.TokoRepository) TokoService {
	return &service{repo: repo}
}

func (s *service) CreateToko(ctx context.Context, data TokoRequest) error {
	toko := &models.Toko{
		UserID:  data.UserID,
		Name:    data.Name,
		Address: data.Address,
	}
	return s.repo.CreateToko(ctx, toko)
}

func (s *service) GetMyToko(ctx context.Context, userID int64) ([]*models.Toko, error) {
	return s.repo.GetMyToko(ctx, userID)
}

func (s *service) UpdateToko(ctx context.Context, data TokoRequest) error {
	// verify the toko exists and belongs to the user
	tokos, err := s.repo.GetMyToko(ctx, data.UserID)
	if err != nil {
		return err
	}

	// Check if the toko with the given ID belongs to the user
	var found bool
	for _, t := range tokos {
		if t.ID == data.ID {
			found = true
			break
		}
	}

	if !found {
		return errors.New("toko not found or you don't have permission to update it")
	}

	toko := &models.Toko{
		ID:      data.ID,
		UserID:  data.UserID,
		Name:    data.Name,
		Address: data.Address,
	}
	return s.repo.UpdateToko(ctx, toko)
}
