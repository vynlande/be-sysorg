package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vynlande/be-sysorg/internal/entity"
	"github.com/vynlande/be-sysorg/pkg/util"
)

func (s *ServiceServer) DeleteExample(c *fiber.Ctx) error {
	s.exampleService.DeleteExample(entity.DeleteExampleReq{
		ID: util.NewQueryHelper().GetIDbyUUID("example", c.Params("id")),
	})
	return c.JSON(util.Response{
		Message: util.MESSAGE_OK_DELETE,
	})
}
