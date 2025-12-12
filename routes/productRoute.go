package routes

import (
	"go-fiber-modular/modules/product/http"
	"go-fiber-modular/modules/product/repository"
	"go-fiber-modular/modules/product/service"
	tokoRepo "go-fiber-modular/modules/toko/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func ProductRouter(app *fiber.App, db *gorm.DB) {
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo)
	tokoRepository := tokoRepo.NewTokoRepository(db)
	productHandler := http.NewProductHandler(productService, tokoRepository)

	http.ProductRoutes(app, productHandler)
}
