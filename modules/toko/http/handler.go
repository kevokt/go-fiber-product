package http

import (
	"go-fiber-modular/helper"
	"go-fiber-modular/modules/toko/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type TokoHandler struct {
	service service.TokoService
}

func NewTokoHandler(service service.TokoService) *TokoHandler {
	return &TokoHandler{service: service}
}

func (h *TokoHandler) CreateToko(c *fiber.Ctx) error {
	ctx := c.UserContext()

	// Get user ID from context
	userIDInterface := c.Locals("user_id")

	// Check if user ID is present in context
	if userIDInterface == nil {
		return c.Status(http.StatusUnauthorized).JSON(helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", "User ID not found in context"))
	}

	// Type assertion check
	userID, ok := userIDInterface.(int64)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(helper.APIResponse("Internal error", http.StatusInternalServerError, "error", "Invalid user ID type"))
	}

	// Bind request body
	req := &CreateTokoRequest{}
	if err := req.bind(c); err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Validation error", http.StatusBadRequest, "error", err.Error()))
	}

	// Create toko
	err := h.service.CreateToko(ctx, service.TokoRequest{
		Name:    req.Name,
		Address: req.Address,
		UserID:  userID,
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helper.APIResponse("Failed to create toko", http.StatusInternalServerError, "error", err.Error()))
	}

	// Return response
	return c.Status(http.StatusCreated).JSON(helper.APIResponse("Toko created successfully", http.StatusCreated, "success", nil))
}

func (h *TokoHandler) GetMyToko(c *fiber.Ctx) error {
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

	// Get toko list by user ID
	tokos, err := h.service.GetMyToko(ctx, userID)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helper.APIResponse("Failed to get toko", http.StatusInternalServerError, "error", err.Error()))
	}
	resp := &GetMyTokoResponse{
		Tokos: tokos,
	}

	return c.Status(http.StatusOK).JSON(helper.APIResponse("Success", http.StatusOK, "success", resp))
}

func (h *TokoHandler) UpdateToko(c *fiber.Ctx) error {
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

	// Get toko ID from URL parameter
	tokoID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Invalid toko ID", http.StatusBadRequest, "error", err.Error()))
	}

	req := &UpdateTokoRequest{}
	if err := req.bind(c); err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Validation error", http.StatusBadRequest, "error", err.Error()))
	}

	err = h.service.UpdateToko(ctx, service.TokoRequest{
		ID:      int64(tokoID),
		Name:    req.Name,
		Address: req.Address,
		UserID:  userID,
	})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helper.APIResponse("Failed to update toko", http.StatusInternalServerError, "error", err.Error()))
	}

	return c.Status(http.StatusOK).JSON(helper.APIResponse("Toko updated successfully", http.StatusOK, "success", nil))
}

/*
Module: Toko
Method Endpoint Function
Name
Keterangan

POST /toko
CreateToko 🔒
Buat Toko baru (UserID diambil otomatis dari Token).

GET /toko/my
GetMyToko 🔒
Tampilkan hanya Toko milik user yang sedang login.

PUT /toko/:id
UpdateToko 🔒
Update data toko.
*/
