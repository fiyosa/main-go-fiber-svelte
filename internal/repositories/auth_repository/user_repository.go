package auth_repository

import (
	"go-fiber-svelte/internal/db"
	"go-fiber-svelte/internal/db/models"
	"go-fiber-svelte/internal/lang"
	auth "go-fiber-svelte/internal/resources/auth_resource"

	"github.com/gofiber/fiber/v2"
)

func UserRepository(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(int)
	database := db.RUN
	var user models.User
	result := database.Preload("Roles").Preload("UserDetails").First(&user, userId)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": lang.T.Convert(lang.T.Get().NOT_FOUND, map[string]any{"operator": lang.T.Get().USER}),
		})
	}
	return c.JSON(fiber.Map{
		"message": lang.T.Convert(lang.T.Get().RETRIEVED_SUCCESSFULLY, map[string]any{"operator": lang.T.Get().USER}),
		"data":    auth.UserToResource(user),
	})
}
