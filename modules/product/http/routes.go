package http

import "github.com/gofiber/fiber/v2"

func ProductRoutes(app *fiber.App, handler *ProductHandler) {
	app.Get("/products/:id", handler.Show)
	app.Post("/products", handler.Store)
	app.Put("/products/:id", handler.Update)
	app.Delete("/products/:id", handler.Delete)
}
