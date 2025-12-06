package http

import (
	"url_shotter/internal/http/handler"
	"url_shotter/internal/http/repo"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(router *fiber.App, urlRepo repo.URLRepository) {
	router.Get("health", handler.HealthHandler)

	// Rotas de URLs
	router.Post("/api/urls", handler.CreateURLHandler(urlRepo))
	router.Get("/api/urls/:id", handler.GetURLHandler(urlRepo))
	router.Get("/:shortCode", handler.RedirectURLHandler(urlRepo))
}

func NewServer(db *gorm.DB) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: "URL Shortener",
	})

	urlRepo := repo.NewURLRepository(db)

	SetupRoutes(app, urlRepo)
	return app
}
