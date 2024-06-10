package main

import (
	"github.com/alexohneander/info-service/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Initialize standard Go html template engine
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	// Configure Middlewares
	app.Use(logger.New())
	app.Use(compress.New())
	// app.Use(cache.New())
	app.Use(healthcheck.New())

	app.Get("/", handler.Default)

	// Metrics
	app.Get("/metrics", monitor.New(monitor.Config{Title: "MyService Metrics Page", APIOnly: true}))

	app.Listen(":4000")
}
