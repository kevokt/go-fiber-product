package http

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// Register Request/Response
type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (r *RegisterRequest) bind(c *fiber.Ctx) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}
	return validate.Struct(r)
}

type RegisterResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// Login Request/Response
type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (r *LoginRequest) bind(c *fiber.Ctx) error {
	if err := c.BodyParser(r); err != nil {
		return err
	}
	return validate.Struct(r)
}

type LoginResponse struct {
	Token string `json:"token"`
}
