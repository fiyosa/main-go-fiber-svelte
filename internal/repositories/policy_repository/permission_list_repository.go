package policy_repository

import (
	"go-fiber-svelte/internal/db"
	"go-fiber-svelte/internal/db/models"
	"go-fiber-svelte/internal/lang"
	"go-fiber-svelte/internal/resources/policy_resource"

	"github.com/gofiber/fiber/v2"
)

func PermissionListRepository(c *fiber.Ctx) error {
	database := db.GetDB()
	var permissions []models.Permission
	database.Find(&permissions)
	return c.JSON(fiber.Map{
		"message": lang.T.Convert(lang.T.Get().RETRIEVED_SUCCESSFULLY, map[string]any{"operator": "Permission"}),
		"data":    policy_resource.PermissionToResource(permissions),
	})
}
