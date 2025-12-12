package http

import (
	"go-fiber-modular/middleware"

	"github.com/gofiber/fiber/v2"
)

func TokoRoutes(app *fiber.App, handler *TokoHandler) {
	toko := app.Group("/toko")
	toko.Post("/", middleware.AuthMiddleware, handler.CreateToko)
	toko.Get("/my", middleware.AuthMiddleware, handler.GetMyToko)
	toko.Put("/:id", middleware.AuthMiddleware, handler.UpdateToko)
}
