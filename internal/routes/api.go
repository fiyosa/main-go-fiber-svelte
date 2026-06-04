package routes

import (
	"go-fiber-svelte/internal/controllers"
	"go-fiber-svelte/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func RegisterAPI(app *fiber.App) {
	auth := middleware.AuthMiddleware()
	api := app.Group("/api")

	api.Get("/openapi.json", controllers.OpenAPI)
	api.Get("/docs", controllers.Docs)
	api.Get("/guest/ping", controllers.Ping)

	authGrp := api.Group("/auth")
	authGrp.Post("/login", controllers.Login)
	authGrp.Delete("/logout", auth, controllers.Logout)
	authGrp.Get("/user", auth, controllers.User)

	policyGrp := api.Group("/policy", auth)
	policyGrp.Get("/role", controllers.RoleList)
	policyGrp.Get("/permission", controllers.PermissionList)
	policyGrp.Post("/permission", controllers.PermissionStore)
	policyGrp.Delete("/permission/:id", controllers.PermissionDestroy)
}
