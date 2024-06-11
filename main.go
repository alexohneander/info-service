package main

import (
	"github.com/alexohneander/info-service/handler"
	"github.com/alexohneander/info-service/middleware/prometheus"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/favicon"
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
	app.Use(favicon.New())

	prometheus := prometheus.New("info-service")
	prometheus.RegisterAt(app, "/metrics")
	app.Use(prometheus.Middleware)

	// Metrics
	app.Get("/monitor", monitor.New(monitor.Config{Title: "Info-Service Monitoring", APIOnly: false}))

	// Routes
	app.All("/", handler.Default)
	app.All("/ip", handler.IP)
	app.All("/ua", handler.UA)
	app.All("/lang", handler.Lang)

	app.Listen(":4000")
}
