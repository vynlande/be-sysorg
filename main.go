package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/vynlande/be-sysorg/config"
	"github.com/vynlande/be-sysorg/env"
	"github.com/vynlande/be-sysorg/infrastructure"
	"github.com/vynlande/be-sysorg/internal/repository"
	"github.com/vynlande/be-sysorg/internal/route"
	"github.com/vynlande/be-sysorg/internal/service"
)

func init() {
	numCPU := runtime.NumCPU()
	if numCPU <= 1 {
		runtime.GOMAXPROCS(1)
	} else {
		runtime.GOMAXPROCS(numCPU / 2)
	}

	env.LoadEnvironmentFile()
	env.NewEnvironment()

	infrastructure.ConnectSqlDB()
	infrastructure.ConnectGormDB()
}

func main() {
	// register service
	dao := repository.NewDAO(infrastructure.SqlDB, infrastructure.GormDB)
	exampleService := service.NewExampleService(dao)

	// logger
	file, err := os.OpenFile("./log/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("logger: load log file - %v \n", err)
	}
	log.SetOutput(file)

	// fiber
	fiberApp := fiber.New(config.NewFiberConfig())
	fiberApp.Use(recover.New())
	route.NewRoute(fiberApp, exampleService)
	fiberApp.Listen(env.SERVER_HOST + ":" + env.SERVER_PORT)
}
