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

	if err := lib.Validate.Check(c, req); err != nil {
		return err
	}

	println("masuk 1", req.Email)
	var user models.User
	result := db.RUN.Where("email = ?", req.Email).First(&user)
	if result.Error != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": lang.T.Get().AUTH_FAILED,
		})
	}
	println("masuk 2")
	if !lib.Hash.Verify(req.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": lang.T.Get().AUTH_FAILED,
		})
	}
	token, err := lib.Jwt.Create(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": lang.T.Get().SOMETHING_WENT_WRONG,
		})
	}
	authRecord := models.Auth{
		UserID:    user.ID,
		Token:     token,
		Revoke:    false,
		IP:        c.IP(),
		UserAgent: c.Get("User-Agent"),
	}
	db.RUN.Create(&authRecord)

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HTTPOnly: true,
		Secure:   config.APP_Env != "local",
	})
	return c.JSON(fiber.Map{
		"message": lang.T.Convert(lang.T.Get().SAVED_SUCCESSFULLY, map[string]any{"operator": "Login"}),
	})
}
