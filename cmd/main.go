package main

import (
	"github.com/architectv/merchantx-task/pkg/handler"
	"github.com/architectv/merchantx-task/pkg/repository"
	"github.com/architectv/merchantx-task/pkg/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	app := fiber.New()
	app.Use(logger.New())
	handlers.InitRoutes(app)
	app.Listen(":8001")
}
