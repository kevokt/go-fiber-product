package routes

import (
	"go-fiber-modular/modules/post/http"
	"go-fiber-modular/modules/post/repository"
	"go-fiber-modular/modules/post/service"

	"github.com/gofiber/fiber/v2"

	"gorm.io/gorm"
)

func PostRouter(app *fiber.App, db *gorm.DB) {
	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	postHandler := http.NewPostHandler(postService)

	http.PostRoutes(app, postHandler)
}
