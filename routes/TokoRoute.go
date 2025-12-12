package routes

import (
	"go-fiber-modular/modules/toko/http"
	"go-fiber-modular/modules/toko/repository"
	"go-fiber-modular/modules/toko/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TokoRouter(app *fiber.App, db *gorm.DB) {
	tokoRepo := repository.NewTokoRepository(db)
	tokoService := service.NewTokoService(tokoRepo)
	tokoHandler := http.NewTokoHandler(tokoService)

	http.TokoRoutes(app, tokoHandler)
}
