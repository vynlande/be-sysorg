package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vynlande/be-sysorg/pkg/util"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: util.NewError().Handler,
	}
}
