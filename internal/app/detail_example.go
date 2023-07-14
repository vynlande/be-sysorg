package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vynlande/be-sysorg/internal/entity"
	"github.com/vynlande/be-sysorg/pkg/util"
)

func (s *ServiceServer) DetailExample(c *fiber.Ctx) error {
	service := s.exampleService.DetailExample(entity.DetailExampleReq{
		ID: util.NewQueryHelper().GetIDbyUUID("example", c.Params("id")),
	})
	return c.JSON(util.Response{
		Data:    service.Item,
		Message: util.MESSAGE_OK_READ,
	})
}
