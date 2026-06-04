package openapi

import "github.com/gofiber/fiber/v2"

type OpenAPIDoc struct {
	OpenAPI string   `json:"openapi"`
	Info    Info     `json:"info"`
	Paths   []string `json:"paths,omitempty"`
}

type Info struct {
	Title   string `json:"title"`
	Version string `json:"version"`
}

func Generate() *OpenAPIDoc {
	return &OpenAPIDoc{
		OpenAPI: "3.0.0",
		Info: Info{
			Title:   "go-fiber-svelte API",
			Version: "1.0.0",
		},
	}
}

func Serve(c *fiber.Ctx) error {
	return c.JSON(Generate())
}
