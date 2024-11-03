package main

import (
	"log"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pageton/SvelteTelegramApp/config"
	"github.com/pageton/SvelteTelegramApp/handler"
	"github.com/pageton/SvelteTelegramApp/middleware"
)

func main() {
	app := fiber.New(
		fiber.Config{
			JSONEncoder: sonic.Marshal,
			JSONDecoder: sonic.Unmarshal,
		})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "POST, GET, OPTIONS",
	}))

	app.Use(middleware.RequestLogger)
	app.Use(middleware.ErrorHandling)

	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	if err := cfg.Validate(); err != nil {
		log.Fatalf("Invalid configuration: %v", err)
	}

	app.Post("/validate", func(ctx *fiber.Ctx) error {
		return handler.ValidateHashHandler(ctx)
	})

	log.Fatal(app.Listen(cfg.Port))
}
