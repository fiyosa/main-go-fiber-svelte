package policy_repository

import (
	"go-fiber-svelte/internal/db"
	"go-fiber-svelte/internal/db/models"
	"go-fiber-svelte/internal/lang"
	"go-fiber-svelte/internal/resources/policy_resource"

	"github.com/gofiber/fiber/v2"
)

func RoleListRepository(c *fiber.Ctx) error {
	database := db.RUN
	var roles []models.Role
	database.Preload("Permissions").Find(&roles)
	return c.JSON(fiber.Map{
		"message": lang.T.Convert(lang.T.Get().RETRIEVED_SUCCESSFULLY, map[string]any{"operator": lang.T.Get().ROLE}),
		"data":    policy_resource.RoleListToResource(roles),
	})
}
