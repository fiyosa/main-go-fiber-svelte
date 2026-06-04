package auth_repository

import (
	"go-fiber-svelte/internal/config"
	"go-fiber-svelte/internal/db"
	"go-fiber-svelte/internal/db/models"
	"go-fiber-svelte/internal/lang"
	"go-fiber-svelte/internal/lib"
	"go-fiber-svelte/internal/request/auth_request"

	"github.com/gofiber/fiber/v2"
)

func LoginRepository(c *fiber.Ctx) error {
	req := new(auth_request.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": lang.T("invalid_request_body"),
		})
	}
	database := db.GetDB()
	var user models.User
	result := database.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": lang.T("invalid_credentials"),
		})
	}
	if !lib.Hash.Verify(req.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": lang.T("invalid_credentials"),
		})
	}
	token, err := lib.Jwt.Create(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": lang.T("failed_generate_token"),
		})
	}
	authRecord := models.Auth{
		UserID:    user.ID,
		Token:     token,
		Revoke:    false,
		IP:        c.IP(),
		UserAgent: c.Get("User-Agent"),
	}
	database.Create(&authRecord)

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HTTPOnly: true,
		Secure:   config.APP_Env != "local",
	})
	return c.JSON(fiber.Map{
		"message": lang.T("login_successful"),
		"data": fiber.Map{
			"token": token,
			"user":  user,
		},
	})
}
