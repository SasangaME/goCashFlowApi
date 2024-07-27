package routes

import (
	"github.com/SasangaME/goCashFlowApi/pkg/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/build", handler.GetBuild)

	// role
	app.Get("/api/roles", handler.RoleGetAll)
	app.Get("/api/roles/:id", handler.RoleGetById)
	app.Post("/api/roles", handler.RoleCreate)
	app.Put("/api/roles/:id", handler.RoleUpdate)

	// user
	app.Get("/api/users", handler.UserGetAll)
	app.Get("/api/users/:id", handler.UserGetById)
	app.Post("/api/users", handler.UserCreate)

	// auth
	app.Put("/api/auth/login", handler.Login)
}
