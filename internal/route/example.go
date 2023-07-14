package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vynlande/be-sysorg/internal/app"
)

func newExampleRoute(a *fiber.App, handler *app.ServiceServer) {
	route := a.Group("/example")
	route.Post("/", handler.CreateExample)
	route.Get("/", handler.ListExample)
	route.Get("/:id", handler.DetailExample)
	route.Put("/:id", handler.PutExample)
	route.Patch("/:id", handler.PatchExample)
	route.Delete("/:id", handler.DeleteExample)
}
