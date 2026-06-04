package policy_repository

import (
	"go-fiber-svelte/internal/db"
	"go-fiber-svelte/internal/db/models"
	"go-fiber-svelte/internal/lang"
	"go-fiber-svelte/internal/resources/policy_resource"

	"github.com/gofiber/fiber/v2"
)

func RoleListRepository(c *fiber.Ctx) error {
	database := db.GetDB()
	var roles []models.Role
	database.Preload("Permissions").Find(&roles)
	return c.JSON(fiber.Map{
		"message": lang.T("roles_retrieved"),
		"data":    policy_resource.RoleListToResource(roles),
	})
}
