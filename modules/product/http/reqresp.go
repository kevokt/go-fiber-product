package http

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type ProductDetail struct {
	ID          int64     `json:"id"`
	Product     string    `json:"product"`
	Description string    `json:"description"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type (
	ShowRequest struct {
		ID int64 `params:"id"`
	}
	ShowResponse ProductDetail
)

func (r *ShowRequest) bind(c *fiber.Ctx) error {
	return c.ParamsParser(r)
}

type (
	StoreRequest struct {
		Product     string `json:"product"`
		Description string `json:"description"`
		Quantity    int    `json:"quantity"`
	}
	StoreResponse ProductDetail
)

func (r *StoreRequest) bind(c *fiber.Ctx) error {
	return c.BodyParser(r)
}

type (
	UpdateRequest struct {
		ID          int64  `params:"id"`
		Product     string `json:"product"`
		Description string `json:"description"`
		Quantity    int    `json:"quantity"`
	}
	UpdateResponse ProductDetail
)

func (r *UpdateRequest) bind(c *fiber.Ctx) error {
	err := c.ParamsParser(r)
	if err != nil {
		return err
	}
	err = c.BodyParser(r)
	if err != nil {
		return err
	}
	return nil
}

type (
	DeleteRequest struct {
		ID int64 `params:"id"`
	}
	DeleteResponse struct {
		Deleted bool `json:"deleted"`
	}
)

func (r *DeleteRequest) bind(c *fiber.Ctx) error {
	return c.ParamsParser(r)
}
