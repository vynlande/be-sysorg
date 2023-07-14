package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vynlande/be-sysorg/internal/app"
	"github.com/vynlande/be-sysorg/internal/service"
)

func NewRoute(fiberApp *fiber.App, exampleService service.ExampleService) {
	allHandler := app.NewService(exampleService)
	newExampleRoute(fiberApp, allHandler)
}
