package service

import (
	"context"

	"go-fiber-modular/models"
	"go-fiber-modular/modules/product/repository"
)

type service struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &service{repo: repo}
}

func (s *service) FindByID(ctx context.Context, id int64) (*models.Product, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) Create(ctx context.Context, data CreateProductData) (*models.Product, error) {
	product := &models.Product{
		Product:     data.Product,
		Description: data.Description,
		Quantity:    data.Quantity,
	}

	err := s.repo.Create(ctx, product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *service) Update(ctx context.Context, data UpdateProductData) (*models.Product, error) {
	product, err := s.repo.FindByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, err
	}

	product.Product = data.Product
	product.Description = data.Description
	product.Quantity = data.Quantity

	err = s.repo.Update(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *service) DeleteByID(ctx context.Context, id int64) (*models.Product, error) {
	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, err
	}

	err = s.repo.Delete(ctx, product)

	return product, err
}
