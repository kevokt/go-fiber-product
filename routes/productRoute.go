package routes

import (
	"go-fiber-modular/modules/product/http"
	"go-fiber-modular/modules/product/repository"
	"go-fiber-modular/modules/product/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProductRouter(app *fiber.App, db *gorm.DB) {
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	productHandler := http.NewProductHandler(productService)

	http.ProductRoutes(app, productHandler)
}
