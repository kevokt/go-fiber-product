package http

import (
	"go-fiber-modular/helper"
	"go-fiber-modular/models"
	"go-fiber-modular/modules/product/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	service service.ProductService
}

func NewProductHandler(service service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) Show(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req := &ShowRequest{}
	if err := req.bind(c); err != nil {
		return err
	}

	product, err := h.service.FindByID(ctx, req.ID)
	if err != nil {
		return err
	}
	// ensure product exists or return 404
	if product == nil {
		c.Status(http.StatusNotFound).JSON(helper.APIResponse("Product is not found", http.StatusNotFound, "Error", []models.Product{}))
		return nil
	}

	resp := &ShowResponse{
		ID:          product.ID,
		Product:     product.Product,
		Description: product.Description,
		Quantity:    product.Quantity,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}

	c.Status(http.StatusOK).JSON(helper.APIResponse("Success Get Product", http.StatusOK, "Success", resp))
	return nil
}

func (h *ProductHandler) Store(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req := &StoreRequest{}
	if err := req.bind(c); err != nil {
		c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Failed to Store Product", http.StatusBadRequest, "Error", err))
		return nil
	}

	product, err := h.service.Create(ctx, service.CreateProductData{
		Product:     req.Product,
		Description: req.Description,
		Quantity:    req.Quantity,
	})
	if err != nil {
		return err
	}

	resp := &StoreResponse{
		ID:          product.ID,
		Product:     product.Product,
		Description: product.Description,
		Quantity:    product.Quantity,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}

	c.Status(http.StatusOK).JSON(helper.APIResponse("Success Store Product", http.StatusOK, "Success", resp))
	return nil
}

func (h *ProductHandler) Update(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req := &UpdateRequest{}
	if err := req.bind(c); err != nil {
		c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Failed to Update Product", http.StatusBadRequest, "Error", err))
		return nil
	}

	product, err := h.service.Update(ctx, service.UpdateProductData{
		ID:          req.ID,
		Product:     req.Product,
		Description: req.Description,
		Quantity:    req.Quantity,
	})
	if err != nil {
		return err
	}
	// ensure product exists or return 404
	if product == nil {
		c.Status(http.StatusNotFound).JSON(helper.APIResponse("Product is not found", http.StatusNotFound, "Error", []models.Product{}))
		return nil
	}

	resp := &UpdateResponse{
		ID:          product.ID,
		Product:     product.Product,
		Description: product.Description,
		Quantity:    product.Quantity,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}

	c.Status(http.StatusOK).JSON(helper.APIResponse("Success Update Product", http.StatusOK, "Success", resp))
	return nil
}

func (h *ProductHandler) Delete(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req := &DeleteRequest{}
	if err := req.bind(c); err != nil {
		c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Failed to Delete Product", http.StatusBadRequest, "Error", err))
		return nil
	}

	product, err := h.service.DeleteByID(ctx, req.ID)
	if err != nil {
		return err
	}
	// ensure product exists or return 404
	if product == nil {
		c.Status(http.StatusNotFound).JSON(helper.APIResponse("Product is not found", http.StatusNotFound, "Error", []models.Product{}))
		return nil
	}

	c.Status(http.StatusOK).JSON(helper.APIResponse("Success Delete Product", http.StatusOK, "Success", []models.Product{}))
	return nil
}
