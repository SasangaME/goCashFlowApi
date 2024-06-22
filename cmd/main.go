package main

import (
	"github.com/SasangaME/goCashFlowApi/pkg/routes"
	"github.com/SasangaME/goCashFlowApi/pkg/util"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	routes.SetupRoutes(app)

	port, err := util.Config("PORT")
	if err != nil {
		log.Fatal(err)
	}

	err = app.Listen(port)
	if err != nil {
		panic(err)
	}
}
