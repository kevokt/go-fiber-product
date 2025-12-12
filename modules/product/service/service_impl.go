package service

import (
	"context"
	"errors"
	"go-fiber-modular/models"
	"go-fiber-modular/modules/product/repository"
)

type service struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &service{repo: repo}
}

func (s *service) CreateProduct(ctx context.Context, data ProductRequest) error {
	product := &models.Product{
		TokoID:      data.TokoID,
		Product:     data.Product,
		Description: data.Description,
		Quantity:    data.Quantity,
	}
	return s.repo.CreateProduct(ctx, product)
}

func (s *service) ListByToko(ctx context.Context, tokoID int64) ([]*models.Product, error) {
	return s.repo.ListByToko(ctx, tokoID)
}

func (s *service) GetProduct(ctx context.Context, productID int64) (*models.Product, error) {
	return s.repo.GetProduct(ctx, productID)
}

func (s *service) UpdateProduct(ctx context.Context, data ProductRequest) error {
	// verify the product exists and belongs to the toko
	product, err := s.repo.GetProduct(ctx, data.ID)
	if err != nil {
		return err
	}

	if product.TokoID != data.TokoID {
		return errors.New("product not found or you don't have permission to update it")
	}

	product.Product = data.Product
	product.Description = data.Description
	product.Quantity = data.Quantity
	return s.repo.UpdateProduct(ctx, product)
}

func (s *service) DeleteProduct(ctx context.Context, productID int64) error {
	return s.repo.DeleteProduct(ctx, productID)
}

/*
POST /products
CreateProduct 🔒
Input Body wajibkirim toko_id. Validasi kepemilikan toko.

GET /toko/:id/products ListByToko 🔒 Tampilkan semua
produk berdasarkan ID
Toko tertentu.

GET /products/:id GetProduct 🔒 Lihat detail satu
produk.

PUT /products/:id UpdateProduct 🔒 Update produk
(Cek apakah user
pemilik toko produk
tsb).

DELETE /products/:id DeleteProduct 🔒 Hapus produk (Cek
apakah user pemilik
toko produk tsb)
*/
