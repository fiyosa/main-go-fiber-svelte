package logger_repository

import (
	"os"
	"strings"

	"go-fiber-svelte/internal/helper"

	"github.com/gofiber/fiber/v2"
)

func LogDeleteRepository(c *fiber.Ctx) error {
	filename := c.Params("filename")
	if filename == "" || strings.Contains(filename, "..") {
		return c.Status(fiber.StatusBadRequest).JSON(helper.Res.Error("Invalid filename", nil))
	}

	if err := os.Remove("./logs/" + filename); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(helper.Res.Error("Log file not found", nil))
	}

	return c.JSON(helper.Res.Success("Log file deleted successfully"))
}
