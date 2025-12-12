package routes

import (
	"go-fiber-modular/modules/auth/http"
	"go-fiber-modular/modules/auth/repository"
	"go-fiber-modular/modules/auth/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRouter(app *fiber.App, db *gorm.DB) {
	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authHandler := http.NewAuthHandler(authService)

	http.AuthRoutes(app, authHandler)
}
