package main

import (
	"github.com/architectv/merchantx-task/pkg/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	handlers := new(handler.Handler)

	app := fiber.New()
	app.Use(logger.New())
	handlers.InitRoutes(app)
	app.Listen(":8001")
}
