package app

import "github.com/vynlande/be-sysorg/internal/service"

type ServiceServer struct {
	exampleService service.ExampleService
}

func NewService(exampleService service.ExampleService) *ServiceServer {
	return &ServiceServer{
		exampleService: exampleService,
	}
}
