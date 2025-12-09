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

	app := fiber.New()
	// Use the cors middleware to allow all origins and methods
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	// Register your routes here
	routes.PostRouter(app, db)
	routes.ProductRouter(app, db)

	log.Fatalln(app.Listen(":" + os.Getenv("PORT")))
}
