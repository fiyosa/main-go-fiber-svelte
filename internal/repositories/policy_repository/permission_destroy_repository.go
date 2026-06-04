package policy_repository

import (
	"go-fiber-svelte/internal/db"
	"go-fiber-svelte/internal/db/models"
	"go-fiber-svelte/internal/lang"

	"github.com/gofiber/fiber/v2"
)

func PermissionDestroyRepository(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": lang.T("invalid_id"),
		})
	}
	database := db.GetDB()
	result := database.Delete(&models.Permission{}, id)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": lang.T("permission_not_found"),
		})
	}
	return c.JSON(fiber.Map{
		"message": lang.T("permission_deleted"),
	})
}
