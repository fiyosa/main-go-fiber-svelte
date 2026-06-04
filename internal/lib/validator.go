package lib

import (
	"encoding/json"
	"go-fiber-svelte/internal/config"
	"go-fiber-svelte/internal/lang"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var Validate validate

type validate struct{}

func (validate) Check(c *fiber.Ctx, input any) (err error, isOk bool) {
	if err := c.BodyParser(input); err != nil {
		return generateError(c, err), false
	}

	if err := config.Validate.Struct(input); err != nil {
		return generateError(c, err), false
	}

	return nil, true
}

func generateError(c *fiber.Ctx, err error) error {
	newErrors := map[string]string{}
	msg := "Invalid data"

	switch v := err.(type) {
	case *json.UnmarshalTypeError:
		field := strings.ToLower(v.Field)
		newErrors[field] = "Json binding error: " + field + " type error"

	case validator.ValidationErrors:
		for _, e := range v {
			field := strings.ToLower(e.Field())
			newErrors[field] = strings.ToLower(e.Translate(config.Translator))
		}

	default:
		if v != nil {
			msg = v.Error()
		} else {
			msg = lang.T.Get().SOMETHING_WENT_WRONG
		}
	}

	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"message": msg,
		"errors":  newErrors,
	})
}
