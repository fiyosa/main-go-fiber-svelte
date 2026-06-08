package main

import (
	"go-fiber-svelte/internal/bootstrap"
	"go-fiber-svelte/internal/config"
	"go-fiber-svelte/internal/helper"
	"go-fiber-svelte/internal/lib"
	"go-fiber-svelte/internal/provider"
	"go-fiber-svelte/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	bootstrap.Init()

	app := fiber.New(fiber.Config{
		ErrorHandler: provider.NewErrorHandler(),
	})

	app.Use(recover.New())
	app.Use(fiberLogger.New())
	app.Use(cors.New(config.CORSConfig()))

	provider.RegisterMiddleware(app)
	routes.RegisterAPI(app)

	app.Static("/", "public")
	app.Get("/api/*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(helper.Res.Error("API endpoint not found", nil))
	})
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("public/index.html")
	})

	port := config.APP_PORT
	if port == "" {
		port = "8000"
	}
	lib.Log.Info("Server started", "fiber", ":"+port)
	app.Listen(":" + port)
}
