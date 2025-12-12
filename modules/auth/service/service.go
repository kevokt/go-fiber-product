package service

import (
	"context"
	"go-fiber-modular/models"
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

type AuthService interface {
	Register(ctx context.Context, data RegisterUserData) error
	Login(ctx context.Context, data LoginUserData) (*models.User, error)
}

type RegisterUserData struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type LoginUserData struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
