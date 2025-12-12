package http

import (
	"go-fiber-modular/helper"
	"go-fiber-modular/middleware"
	"go-fiber-modular/modules/auth/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

/*
Method Endpoint Function
Name
Keterangan

POST /auth/register Register
Input: Name, Email, Password.
Output: Created User.

POST /auth/login Login
Input: Email, Password.
Output: JWT Token
*/

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(service service.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req := &RegisterRequest{}
	if err := req.bind(c); err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Validation error", http.StatusBadRequest, "error", err.Error()))
	}

	err := h.service.Register(ctx, service.RegisterUserData{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helper.APIResponse("Failed to register user", http.StatusInternalServerError, "error", err.Error()))
	}

	return c.Status(http.StatusCreated).JSON(helper.APIResponse("User registered successfully", http.StatusCreated, "success", nil))
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req := &LoginRequest{}
	if err := req.bind(c); err != nil {
		return c.Status(http.StatusBadRequest).JSON(helper.APIResponse("Validation error", http.StatusBadRequest, "error", err.Error()))
	}

	user, err := h.service.Login(ctx, service.LoginUserData{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(helper.APIResponse("Invalid email or password", http.StatusUnauthorized, "error", nil))
	}

	token, err := middleware.GenerateToken(user.ID, user.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(helper.APIResponse("Failed to generate token", http.StatusInternalServerError, "error", nil))
	}

	resp := &LoginResponse{
		Token: token,
	}

	return c.Status(http.StatusOK).JSON(helper.APIResponse("Login successful", http.StatusOK, "success", resp))
}
