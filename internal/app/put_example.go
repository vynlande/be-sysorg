package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vynlande/be-sysorg/internal/entity"
	"github.com/vynlande/be-sysorg/pkg/http/request"
	"github.com/vynlande/be-sysorg/pkg/util"
)

func (s *ServiceServer) PutExample(c *fiber.Ctx) error {
	body := new(request.PutExample)
	err := c.BodyParser(&body)
	util.NewError().Panic(err)
	s.exampleService.PutExample(entity.PutExampleReq{
		ID:   util.NewQueryHelper().GetIDbyUUID("example", c.Params("id")),
		Name: body.Name,
		Age:  body.Age,
	})
	return c.JSON(util.Response{
		Message: util.MESSAGE_OK_UPDATE,
	})
}
