package http

import (
	"go-fiber-modular/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// Create Product Request/Response
type CreateProductRequest struct {
	TokoID      int64  `grom:"index" json:"toko_id" validate:"required"`
	Product     string `json:"product" validate:"required"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity" validate:"required"`
}

type CreateProductResponse struct {
	ID int64 `json:"id"`
}

func (r *CreateProductRequest) bind(c *fiber.Ctx) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}
	return validate.Struct(r)
}

// List By Toko Request/Response
type ListByTokoResponse struct {
	Products []*models.Product `json:"products"`
}

// Get Product Request/Response
type GetProductResponse struct {
	Product *models.Product `json:"product"`
}

// Update Product Request/Response
type UpdateProductRequest struct {
	ID          int64  `json:"id"`
	TokoID      int64  `json:"toko_id" validate:"required"`
	Product     string `json:"product" validate:"required"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity" validate:"required"`
}

func (r *UpdateProductRequest) bind(c *fiber.Ctx) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}
	return validate.Struct(r)
}
