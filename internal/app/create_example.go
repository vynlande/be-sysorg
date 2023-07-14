package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vynlande/be-sysorg/internal/entity"
	"github.com/vynlande/be-sysorg/pkg/http/request"
	"github.com/vynlande/be-sysorg/pkg/util"
)

func (s *ServiceServer) CreateExample(c *fiber.Ctx) error {
	body := new(request.CreateExample)
	err := c.BodyParser(&body)
	util.NewError().Panic(err)
	s.exampleService.CreateExample(entity.CreateExampleReq{
		Name: body.Name,
		Age:  body.Age,
	})
	return c.JSON(util.Response{
		Message: util.MESSAGE_OK_CREATE,
	})
}
