package policy_repository

import (
	"go-fiber-svelte/internal/db"
	"go-fiber-svelte/internal/db/models"
	"go-fiber-svelte/internal/lang"
	"go-fiber-svelte/internal/request/policy_request"

	"github.com/gofiber/fiber/v2"
)

func PermissionStoreRepository(c *fiber.Ctx) error {
	req := new(policy_request.PermissionStoreRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": lang.T("invalid_request_body"),
		})
	}
	permission := models.Permission{
		Name:  req.Name,
		Notes: req.Notes,
	}
	database := db.GetDB()
	if err := database.Create(&permission).Error; err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": lang.T("permission_already_exists"),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": lang.T("permission_created"),
		"data":    permission,
	})
}
