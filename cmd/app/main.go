package main

import (
	"strings"

	"go-fiber-svelte/internal/config"
	"go-fiber-svelte/internal/db"
	"go-fiber-svelte/internal/provider"
	"go-fiber-svelte/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	config.InitConfigApp()
	db.Init()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": err.Error(),
			})
		},
	})

	app.Use(recover.New())
	app.Use(fiberLogger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: strings.Join([]string{fiber.MethodGet, fiber.MethodPost, fiber.MethodPut, fiber.MethodDelete, fiber.MethodOptions}, ","),
		AllowHeaders: "Content-Type, Authorization",
	}))

	provider.RegisterMiddleware(app)
	routes.RegisterAPI(app)

	app.Static("/", "public")
	app.Get("/api/*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "API endpoint not found",
		})
	})
	app.Get("/*", func(c *fiber.Ctx) error {
		return c.SendFile("public/index.html")
	})

	port := config.APP_Port
	if port == "" {
		port = "8000"
	}
	app.Listen(":" + port)
}
