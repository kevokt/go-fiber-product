package repository

import (
	"context"
	"go-fiber-modular/models"

	"gorm.io/gorm"
)

type tokoRepository struct {
	db *gorm.DB
}

func NewTokoRepository(db *gorm.DB) TokoRepository {
	return &tokoRepository{db: db}
}

func (r *tokoRepository) CreateToko(ctx context.Context, toko *models.Toko) error {
	return r.db.Create(toko).Error
}

func (r *tokoRepository) GetMyToko(ctx context.Context, userID int64) ([]*models.Toko, error) {
	var tokos []*models.Toko
	err := r.db.Where("user_id = ?", userID).Find(&tokos).Error
	if err != nil {
		return nil, err
	}
	return tokos, nil
}

func (r *tokoRepository) GetTokoByID(ctx context.Context, tokoID int64) (*models.Toko, error) {
	var toko models.Toko
	err := r.db.Where("id = ?", tokoID).First(&toko).Error
	if err != nil {
		return nil, err
	}
	return &toko, nil
}

func (r *tokoRepository) UpdateToko(ctx context.Context, toko *models.Toko) error {
	return r.db.Model(&models.Toko{}).Where("id = ?", toko.ID).Updates(map[string]any{
		"name":    toko.Name,
		"address": toko.Address,
	}).Error
}
