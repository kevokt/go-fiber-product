package config

import (
	"go-fiber-modular/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"gorm.io/gorm"
)

func Route(db *gorm.DB) {

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			// Return JSON error response
			return c.Status(code).JSON(fiber.Map{
				"meta": fiber.Map{
					"message": "Error",
					"code":    code,
					"status":  "error",
				},
				"data": err.Error(),
			})
		},
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	// Register your routes here
	routes.UserRouter(app, db)
	routes.TokoRouter(app, db)
	routes.ProductRouter(app, db)

	log.Fatalln(app.Listen(":" + os.Getenv("PORT")))
}
