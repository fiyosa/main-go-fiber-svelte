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
			"message": "Invalid ID",
		})
	}
	database := db.GetDB()
	result := database.Delete(&models.Permission{}, id)
	if result.RowsAffected == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": lang.T.Convert(lang.T.Get().NOT_FOUND, map[string]any{"operator": "Permission"}),
		})
	}
	return c.JSON(fiber.Map{
		"message": lang.T.Convert(lang.T.Get().DELETED_SUCCESSFULLY, map[string]any{"operator": "Permission"}),
	})
}
