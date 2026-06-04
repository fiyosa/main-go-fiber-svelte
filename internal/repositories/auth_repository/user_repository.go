package auth_repository

import (
	"go-fiber-svelte/internal/db"
	"go-fiber-svelte/internal/db/models"
	"go-fiber-svelte/internal/lang"
	"go-fiber-svelte/internal/resources/auth_resource"

	"github.com/gofiber/fiber/v2"
)

func UserRepository(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(int)
	database := db.GetDB()
	var user models.User
	result := database.Preload("Roles").Preload("UserDetails").First(&user, userId)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": lang.T("user_not_found"),
		})
	}
	return c.JSON(fiber.Map{
		"message": lang.T("user_retrieved"),
		"data":    auth.UserToResource(user),
	})
}
