package handler

import (
	"url_shotter/internal/http/domain"
	"url_shotter/internal/http/repo"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateURLHandler(repo repo.URLRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req struct {
			OriginalURL string `json:"original_url"`
		}

		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "URL inválida",
			})
		}

		shortCode := uuid.New().String()[:8]

		url := &domain.URL{
			OriginalURL: req.OriginalURL,
			ShortCode:   shortCode,
		}

		if err := repo.Create(url); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Erro ao criar URL",
			})
		}

		return c.Status(fiber.StatusCreated).JSON(url)
	}
}

func GetURLHandler(repo repo.URLRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		url, err := repo.FindByShortCode(id)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "URL não encontrada",
			})
		}

		return c.JSON(url)
	}
}

func RedirectURLHandler(repo repo.URLRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {
		shortCode := c.Params("shortCode")

		url, err := repo.FindByShortCode(shortCode)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "URL não encontrada",
			})
		}

		url.ClickCount++
		repo.Update(url)

		return c.Redirect(url.OriginalURL, fiber.StatusMovedPermanently)
	}
}
