package controllers

import (
	"go-fiber-svelte/internal/http/repositories/logger_repository"

	"github.com/gofiber/fiber/v2"
)

func LoggerList(c *fiber.Ctx) error {
	return logger_repository.LogListRepository(c)
}

func LoggerDetail(c *fiber.Ctx) error {
	return logger_repository.LogDetailRepository(c)
}

func LoggerOpenAPIPaths() map[string]any {
	return map[string]any{
		"/api/guest/logs": map[string]any{
			"get": map[string]any{
				"summary":     "List log files",
				"description": "Returns list of available log files",
				"tags":        []string{"Guest"},
				"responses": map[string]any{
					"200": map[string]any{"description": "List of log files"},
				},
			},
		},
		"/api/guest/logs/{filename}": map[string]any{
			"get": map[string]any{
				"summary":     "Get log detail",
				"description": "Returns log entries from a specific file",
				"tags":        []string{"Guest"},
				"parameters": []map[string]any{
					{"name": "filename", "in": "path", "required": true, "schema": map[string]any{"type": "string"}},
				},
				"responses": map[string]any{
					"200": map[string]any{"description": "Log entries"},
				},
			},
		},
	}
}
