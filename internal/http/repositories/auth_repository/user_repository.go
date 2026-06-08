package auth_repository

import (
	"go-fiber-svelte/internal/db"
	"go-fiber-svelte/internal/db/models"
	"go-fiber-svelte/internal/helper"
	"go-fiber-svelte/internal/http/resources/auth_resource"
	"go-fiber-svelte/internal/lang"

	"github.com/gofiber/fiber/v2"
)

func UserRepository(c *fiber.Ctx) error {
	userId := c.Locals("user_id").(int)
	database := db.RUN
	var user models.User
	result := database.Preload("Roles").Preload("UserDetails").First(&user, userId)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(helper.Res.Error(lang.T.Convert(lang.T.Get().NOT_FOUND, map[string]any{"operator": lang.T.Get().USER}), nil))
	}
	return c.JSON(helper.Res.SuccessData(auth_resource.UserToResource(user), lang.T.Convert(lang.T.Get().RETRIEVED_SUCCESSFULLY, map[string]any{"operator": lang.T.Get().USER})))
}
