package routes

import (
	"github.com/SasangaME/goCashFlowApi/pkg/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/build", handler.GetBuild)
}
