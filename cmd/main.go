package main

import (
	"log"

	"github.com/architectv/merchantx-task/pkg/handler"
	"github.com/architectv/merchantx-task/pkg/repository"
	"github.com/architectv/merchantx-task/pkg/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	app := fiber.New()
	app.Use(logger.New())
	handlers.InitRoutes(app)
	app.Listen(viper.GetString("port"))
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
