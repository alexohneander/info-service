package main

import (
	"github.com/alexohneander/info-service/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(compress.New())
	app.Use(cache.New())

	app.Get("/", handler.Default)

	// Metrics
	app.Get("/metrics", monitor.New())

	app.Listen(":3000")
}
