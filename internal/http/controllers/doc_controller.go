package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type openapiDoc struct {
	OpenAPI    string            `json:"openapi"`
	Info       openapiInfo       `json:"info"`
	Paths      map[string]any    `json:"paths"`
	Components openapiComponents `json:"components"`
}


type openapiInfo struct {
	Title   string `json:"title"`
	Version string `json:"version"`
}

type openapiComponents struct {
	SecuritySchemes openapiSecurityScheme `json:"securitySchemes"`
}

type openapiSecurityScheme struct {
	Bearer openapiBearer `json:"bearer"`
}

type openapiBearer struct {
	Type   string `json:"type"`
	Scheme string `json:"scheme"`
}

func mergeMaps(maps ...map[string]any) map[string]any {
	out := make(map[string]any)
	for _, m := range maps {
		for k, v := range m {
			out[k] = v
		}
	}
	return out
}

func OpenAPI(c *fiber.Ctx) error {
	doc := openapiDoc{
		OpenAPI: "3.0.0",
		Info: openapiInfo{
			Title:   "go-fiber-svelte API",
			Version: "1.0.0",
		},
		Paths: mergeMaps(
			AuthOpenAPIPaths(),
			PolicyOpenAPIPaths(),
			GuestOpenAPIPaths(),
			LoggerOpenAPIPaths(),
		),
		Components: openapiComponents{
			SecuritySchemes: openapiSecurityScheme{
				Bearer: openapiBearer{
					Type:   "http",
					Scheme: "bearer",
				},
			},
		},
	}

	return c.JSON(doc)
}

func Docs(c *fiber.Ctx) error {
	c.Set(fiber.HeaderContentType, "text/html; charset=utf-8")
	return c.SendFile("public/openapi.html")
}
