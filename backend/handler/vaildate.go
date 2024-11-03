package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/pageton/SvelteTelegramApp/config"
	"github.com/pageton/SvelteTelegramApp/services"
)

type HashRequest struct {
	Hash string `json:"hash" validate:"required"`
}

func ValidateHashHandler(c *fiber.Ctx) error {
	var request HashRequest
	if err := c.BodyParser(&request); err != nil {
		log.Println("Error parsing request body:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	if request.Hash == "" {
		log.Println("Hash field is missing")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Missing required field hash",
		})
	}

	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := cfg.Validate(); err != nil {
		log.Fatalf("Invalid configuration: %v", err)
	}

	data, err := services.ParseHash(request.Hash)
	if err != nil {
		log.Println("Error parsing hash:", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid hash format",
		})
	}

	isValid, err := services.IsHashValid(data, cfg.BotToken)
	if err != nil {
		log.Println("Error validating hash:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal server error",
		})
	} else if isValid {
		log.Println("Hash is valid")
		return c.JSON(fiber.Map{"ok": true})
	}

	log.Println("Hash is invalid")
	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
		"error": "Invalid hash",
	})
}
