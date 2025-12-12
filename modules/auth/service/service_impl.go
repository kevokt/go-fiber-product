package service

import (
	"context"
	"errors"
	"go-fiber-modular/helper"
	"go-fiber-modular/models"
	"go-fiber-modular/modules/auth/repository"
)

type service struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &service{repo: repo}
}

func (s *service) Register(ctx context.Context, data RegisterUserData) error {
	hashedPassword, err := helper.HashPassword(data.Password)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:     data.Name,
		Email:    data.Email,
		Password: hashedPassword,
	}
	return s.repo.Register(ctx, user)
}

func (s *service) Login(ctx context.Context, data LoginUserData) (*models.User, error) {
	if data.Email == "" || data.Password == "" {
		return nil, errors.New("email or password is empty")
	}

	user, err := s.repo.FindByEmail(ctx, data.Email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if !helper.CheckPasswordHash(data.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
