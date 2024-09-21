package handler

import (
	"github.com/SasangaME/goCashFlowApi/pkg/model/dto"
	"github.com/SasangaME/goCashFlowApi/pkg/util"
	"github.com/gofiber/fiber/v2"
)

func GetBuild(c *fiber.Ctx) error {

	env, _ := util.Config("BUILD_ENV")
	version, _ := util.Config("BUILD_VERSION")
	buildNumebr, _ := util.Config("BUILD_NUMBER")

	build := dto.BuildDto{
		Env:     env,
		Version: version,
		Build:   buildNumebr,
	}
	return c.Status(200).JSON(build)
}
