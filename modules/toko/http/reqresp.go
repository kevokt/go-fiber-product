package http

import (
	"go-fiber-modular/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// Create Toko Request/Response
type CreateTokoRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
}

func (r *CreateTokoRequest) bind(c *fiber.Ctx) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}
	return validate.Struct(r)
}

type CreateTokoResponse struct {
	ID int64 `json:"id"`
}

// Get My Toko Request/Response
type GetMyTokoResponse struct {
	Tokos []*models.Toko `json:"tokos"`
}

// Update Toko Request/Response
type UpdateTokoRequest struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
}

func (r *UpdateTokoRequest) bind(c *fiber.Ctx) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}
	return validate.Struct(r)
}
