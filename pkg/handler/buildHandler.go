package handler

import (
	"github.com/SasangaME/goCashFlowApi/pkg/model/dto"
	"github.com/gofiber/fiber/v2"
)

func GetBuild(c *fiber.Ctx) error {
	build := dto.BuildDto{
		Env:     "dev",
		Version: "0.0.1",
		Build:   "2024062301",
	}
	return c.Status(200).JSON(build)
}
