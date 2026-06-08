package routes

import (
	"go-fiber-svelte/internal/http/controllers"
	"go-fiber-svelte/internal/middleware"
	"go-fiber-svelte/internal/provider"

	"github.com/gofiber/fiber/v2"
)

func RegisterAPI(app *fiber.App) {
	auth := middleware.AuthMiddleware()
	policy := provider.Policy
	api := app.Group("/api")

	api.Get("/openapi.json", controllers.OpenAPI)
	api.Get("/docs", controllers.Docs)
	api.Get("/guest/ping", controllers.Ping)

	authGrp := api.Group("/auth")
	authGrp.Post("/login", controllers.Login)
	authGrp.Delete("/logout", auth, controllers.Logout)
	authGrp.Get("/user", auth, controllers.User)

	policyGrp := api.Group("/policy", auth)
	policyGrp.Get("/role", policy("role-list"), controllers.RoleList)
	policyGrp.Get("/permission", policy("permission-list"), controllers.PermissionList)
	policyGrp.Post("/permission", policy("permission-store"), controllers.PermissionStore)
	policyGrp.Delete("/permission/:id", policy("permission-destroy"), controllers.PermissionDestroy)
}
