package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vynlande/be-sysorg/internal/entity"
	"github.com/vynlande/be-sysorg/pkg/http/request"
	"github.com/vynlande/be-sysorg/pkg/util"
)

func (s *ServiceServer) ListExample(c *fiber.Ctx) error {
	queryParam := new(request.ListExample).Init()
	err := c.QueryParser(&queryParam)
	util.NewError().Panic(err)
	service := s.exampleService.ListExample(entity.ListExampleReq{
		Search: queryParam.Search,
		Page:   queryParam.Page,
		Limit:  queryParam.Limit,
		Order:  queryParam.Order,
	})
	return c.JSON(util.Response{
		Data: service.Items,
		Meta: util.Meta{
			Page:       service.Page,
			Limit:      service.Limit,
			TotalRows:  service.TotalRows,
			TotalPages: service.TotalPages,
		},
		Message: util.MESSAGE_OK_READ,
	})
}
