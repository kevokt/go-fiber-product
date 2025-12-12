package http

import (
	"go-fiber-modular/helper"
	"go-fiber-modular/modules/product/service"
	tokoRepo "go-fiber-modular/modules/toko/repository"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	service  service.ProductService
	tokoRepo tokoRepo.TokoRepository
}

func NewProductHandler(service service.ProductService, tokoRepo tokoRepo.TokoRepository) *ProductHandler {
	return &ProductHandler{
		service:  service,
		tokoRepo: tokoRepo,
	}
}

func (h *ProductHandler) CreateProduct(c *fiber.Ctx) error {
	ctx := c.UserContext()

	// Get user ID from context with type assertion check
	userIDInterface := c.Locals("user_id")
	if userIDInterface == nil {
		return c.Status(http.StatusUnauthorized).JSON(helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", "User ID not found in context"))
	}
	userID, ok := userIDInterface.(int64)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(helper.APIResponse("Internal error", http.StatusInternalServerError, "error", "Invalid user ID type"))
	}

	// Bind request body
	req := &CreateProductRequest{}
	if err := req.bind(c); err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Validation error", http.StatusBadRequest, "error", err.Error()))
	}

	// Verify toko ownership
	toko, err := h.tokoRepo.GetTokoByID(ctx, req.TokoID)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(helper.APIResponse("Toko not found", http.StatusNotFound, "error", "The specified toko does not exist"))
	}
	if toko.UserID != userID {
		return c.Status(http.StatusForbidden).JSON(helper.APIResponse("Forbidden", http.StatusForbidden, "error", "You don't have permission to add products to this toko"))
	}

	// Create product
	err = h.service.CreateProduct(ctx, service.ProductRequest{
		TokoID:      req.TokoID,
		Product:     req.Product,
		Description: req.Description,
		Quantity:    req.Quantity,
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helper.APIResponse("Failed to create product", http.StatusInternalServerError, "error", err.Error()))
	}

	// Return response
	return c.Status(http.StatusCreated).JSON(helper.APIResponse("Product created successfully", http.StatusCreated, "success", nil))
}

func (h *ProductHandler) ListByToko(c *fiber.Ctx) error {
	ctx := c.UserContext()

	// Get toko ID from URL parameter
	tokoID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Invalid toko ID", http.StatusBadRequest, "error", err.Error()))
	}

	// Check if toko exists
	toko, err := h.tokoRepo.GetTokoByID(ctx, int64(tokoID))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(helper.APIResponse("Toko not found", http.StatusNotFound, "error", "The specified toko does not exist"))
	}
	if toko == nil {
		return c.Status(http.StatusNotFound).JSON(helper.APIResponse("Toko not found", http.StatusNotFound, "error", "The specified toko does not exist"))
	}

	products, err := h.service.ListByToko(ctx, int64(tokoID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helper.APIResponse("Failed to get products", http.StatusInternalServerError, "error", err.Error()))
	}

	resp := &ListByTokoResponse{
		Products: products,
	}

	return c.Status(http.StatusOK).JSON(helper.APIResponse("Success", http.StatusOK, "success", resp))
}

func (h *ProductHandler) GetProduct(c *fiber.Ctx) error {
	ctx := c.UserContext()

	// Get product ID from URL parameter
	productID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Invalid product ID", http.StatusBadRequest, "error", err.Error()))
	}

	product, err := h.service.GetProduct(ctx, int64(productID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helper.APIResponse("Failed to get product", http.StatusInternalServerError, "error", err.Error()))
	}
	if product == nil {
		return c.Status(http.StatusNotFound).JSON(helper.APIResponse("Product not found", http.StatusNotFound, "error", "The specified product does not exist"))
	}

	resp := &GetProductResponse{
		Product: product,
	}

	return c.Status(http.StatusOK).JSON(helper.APIResponse("Success", http.StatusOK, "success", resp))
}

func (h *ProductHandler) UpdateProduct(c *fiber.Ctx) error {
	ctx := c.UserContext()

	// Get user ID from context with type assertion check
	userIDInterface := c.Locals("user_id")
	if userIDInterface == nil {
		return c.Status(http.StatusUnauthorized).JSON(helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", "User ID not found in context"))
	}
	userID, ok := userIDInterface.(int64)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(helper.APIResponse("Internal error", http.StatusInternalServerError, "error", "Invalid user ID type"))
	}

	// Get product ID from URL parameter
	productID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Invalid product ID", http.StatusBadRequest, "error", err.Error()))
	}

	// Bind request body
	req := &UpdateProductRequest{}
	if err := req.bind(c); err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Validation error", http.StatusBadRequest, "error", err.Error()))
	}

	// Verify toko ownership
	toko, err := h.tokoRepo.GetTokoByID(ctx, req.TokoID)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(helper.APIResponse("Toko not found", http.StatusNotFound, "error", "The specified toko does not exist"))
	}
	if toko.UserID != userID {
		return c.Status(http.StatusForbidden).JSON(helper.APIResponse("Forbidden", http.StatusForbidden, "error", "You don't have permission to update products in this toko"))
	}

	// Update product
	err = h.service.UpdateProduct(ctx, service.ProductRequest{
		ID:          int64(productID),
		TokoID:      req.TokoID,
		Product:     req.Product,
		Description: req.Description,
		Quantity:    req.Quantity,
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helper.APIResponse("Failed to update product", http.StatusInternalServerError, "error", err.Error()))
	}

	// Return response
	return c.Status(http.StatusOK).JSON(helper.APIResponse("Product updated successfully", http.StatusOK, "success", nil))
}

func (h *ProductHandler) DeleteProduct(c *fiber.Ctx) error {
	ctx := c.UserContext()

	// Get user ID from context with type assertion check
	userIDInterface := c.Locals("user_id")
	if userIDInterface == nil {
		return c.Status(http.StatusUnauthorized).JSON(helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", "User ID not found in context"))
	}
	userID, ok := userIDInterface.(int64)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(helper.APIResponse("Internal error", http.StatusInternalServerError, "error", "Invalid user ID type"))
	}

	// Get product ID from URL parameter
	productID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Invalid product ID", http.StatusBadRequest, "error", err.Error()))
	}

	// Get the product to verify toko ownership
	product, err := h.service.GetProduct(ctx, int64(productID))
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(helper.APIResponse("Product not found", http.StatusNotFound, "error", "The specified product does not exist"))
	}

	// Verify toko ownership
	toko, err := h.tokoRepo.GetTokoByID(ctx, product.TokoID)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(helper.APIResponse("Toko not found", http.StatusNotFound, "error", "The associated toko does not exist"))
	}
	if toko.UserID != userID {
		return c.Status(http.StatusForbidden).JSON(helper.APIResponse("Forbidden", http.StatusForbidden, "error", "You don't have permission to delete this product"))
	}

	// Delete product
	err = h.service.DeleteProduct(ctx, int64(productID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helper.APIResponse("Failed to delete product", http.StatusInternalServerError, "error", err.Error()))
	}

	// Return response
	return c.Status(http.StatusOK).JSON(helper.APIResponse("Product deleted successfully", http.StatusOK, "success", nil))
}
