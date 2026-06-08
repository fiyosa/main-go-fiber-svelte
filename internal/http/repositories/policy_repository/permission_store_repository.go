package policy_repository

import (
	"go-fiber-svelte/internal/db"
	"go-fiber-svelte/internal/db/models"
	"go-fiber-svelte/internal/http/request/policy_request"
	"go-fiber-svelte/internal/lang"

	"github.com/gofiber/fiber/v2"
)

func PermissionStoreRepository(c *fiber.Ctx) error {
	req := new(policy_request.PermissionStoreRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}
	permission := models.Permission{
		Name:  req.Name,
		Notes: req.Notes,
	}
	database := db.RUN
	if err := database.Create(&permission).Error; err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": lang.T.Convert(lang.T.Get().ALREADY_EXIST, map[string]any{"operator": "Permission"}),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": lang.T.Convert(lang.T.Get().SAVED_SUCCESSFULLY, map[string]any{"operator": "Permission"}),
		"data":    permission,
	})
}
