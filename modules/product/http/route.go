package http

import (
	"go-fiber-modular/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App, handler *ProductHandler) {
	app.Post("/products/", middleware.AuthMiddleware, handler.CreateProduct)
	app.Get("/toko/:id/products", middleware.AuthMiddleware, handler.ListByToko)
	app.Get("/products/:id", middleware.AuthMiddleware, handler.GetProduct)
	app.Put("/products/:id", middleware.AuthMiddleware, handler.UpdateProduct)
	app.Delete("/products/:id", middleware.AuthMiddleware, handler.DeleteProduct)
}
