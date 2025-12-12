package repository

import (
	"context"
	"go-fiber-modular/models"

	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(ctx context.Context, product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) DeleteProduct(ctx context.Context, productID int64) error {
	return r.db.Where("id = ?", productID).Delete(&models.Product{}).Error
}

func (r *productRepository) ListByToko(ctx context.Context, tokoID int64) ([]*models.Product, error) {
	var products []*models.Product
	err := r.db.Where("toko_id = ?", tokoID).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (r *productRepository) GetProduct(ctx context.Context, productID int64) (*models.Product, error) {
	var product models.Product
	err := r.db.Where("id = ?", productID).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) UpdateProduct(ctx context.Context, product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) DeleteByToko(ctx context.Context, tokoID int64) error {
	return r.db.Where("toko_id = ?", tokoID).Delete(&models.Product{}).Error
}
