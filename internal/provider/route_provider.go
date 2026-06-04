package provider

import (
	"github.com/gofiber/fiber/v2"
)

type RouteGroup struct {
	prefix string
	app    *fiber.App
}

func NewRouteGroup(prefix string, app *fiber.App) *RouteGroup {
	return &RouteGroup{prefix: prefix, app: app}
}

func (g *RouteGroup) Get(path string, handler fiber.Handler) fiber.Router {
	return g.app.Get(g.prefix+path, handler)
}

func (g *RouteGroup) Post(path string, handler fiber.Handler) fiber.Router {
	return g.app.Post(g.prefix+path, handler)
}

func (g *RouteGroup) Put(path string, handler fiber.Handler) fiber.Router {
	return g.app.Put(g.prefix+path, handler)
}

func (g *RouteGroup) Delete(path string, handler fiber.Handler) fiber.Router {
	return g.app.Delete(g.prefix+path, handler)
}
